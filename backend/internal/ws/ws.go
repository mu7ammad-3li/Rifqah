package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/rifqah/backend/internal/ball"
	"github.com/rifqah/backend/internal/room"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	userID uuid.UUID
}

type Room struct {
	clients    map[*Client]bool
	turnTimer  *time.Timer
	graceTimer *time.Timer
	isGrace    bool
	mutex      sync.Mutex
}

// Hub needs access to RoomService to verify authorization
type Hub struct {
	rooms   map[string]*Room
	mutex   sync.Mutex
	redis   *redis.Client
	ctx     context.Context
	ballSvc *ball.BallService
	roomSvc *room.RoomService // Added field
}

// Update constructor to take room service
func NewHub(rdb *redis.Client, ballSvc *ball.BallService, roomSvc *room.RoomService) *Hub {
	return &Hub{
		rooms:   make(map[string]*Room),
		redis:   rdb,
		ctx:     context.Background(),
		ballSvc: ballSvc,
		roomSvc: roomSvc,
	}
}

type WSMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type BallStateUpdate struct {
	Type   string   `json:"type"`
	Active string   `json:"active"`
	Queue  []string `json:"queue"`
}

func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomID")
	userIDStr := r.URL.Query().Get("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Valid User ID required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		return
	}

	client := &Client{conn: conn, send: make(chan []byte, 256), userID: userID}
	h.registerClient(roomID, client)

	go client.writePump()
	go h.readPump(roomID, client)
}

func (h *Hub) registerClient(roomID string, client *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomID]
	if !ok {
		room = &Room{clients: make(map[*Client]bool)}
		h.rooms[roomID] = room
		go h.subscribeToRoom(roomID)
	}
	room.mutex.Lock()

	if room.isGrace {
		log.Printf("Attempted join during grace period for room %s", roomID)
		room.mutex.Unlock()
		h.mutex.Unlock()
		client.conn.Close()
		return
	}

	room.clients[client] = true
	room.mutex.Unlock()
	h.mutex.Unlock()

	h.broadcastState(roomID)
}

func (h *Hub) unregisterClient(roomID string, client *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomID]
	if !ok {
		h.mutex.Unlock()
		return
	}
	room.mutex.Lock()
	if _, ok := room.clients[client]; ok {
		delete(room.clients, client)
		close(client.send)

		activeID, _ := h.ballSvc.GetActiveSpeaker(roomID)
		if activeID == client.userID.String() {
			log.Printf("Current speaker %s disconnected, passing the ball", activeID)
			room.mutex.Unlock()
			h.passBall(roomID)
			room.mutex.Lock()
		}
	}

	if len(room.clients) == 0 && !room.isGrace {
		h.enterGracePeriod(roomID, room)
	}
	room.mutex.Unlock()
	h.mutex.Unlock()
}

func (h *Hub) enterGracePeriod(roomID string, room *Room) {
	log.Printf("Entering grace period for room %s", roomID)
	room.isGrace = true

	// Broadcast grace period entry to remaining clients
	h.redis.Publish(h.ctx, "room:"+roomID, `{"type":"GRACE_PERIOD_START"}`)

	room.graceTimer = time.AfterFunc(5*time.Minute, func() {
		h.mutex.Lock()
		defer h.mutex.Unlock()

		h.mutex.Lock()
		room.mutex.Lock()
		log.Printf("Grace period expired for room %s, broadcasting PURGE_ALL", roomID)

		// Unnegotiable final purge
		h.redis.Publish(h.ctx, "room:"+roomID, `{"type":"PURGE_ALL"}`)

		delete(h.rooms, roomID)
		room.mutex.Unlock()
		h.mutex.Unlock()
	})
}

func (h *Hub) readPump(roomID string, client *Client) {
	defer func() {
		h.unregisterClient(roomID, client)
		client.conn.Close()
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			break
		}

		var wsMsg WSMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			h.redis.Publish(h.ctx, "room:"+roomID, message)
			continue
		}

		switch wsMsg.Type {
		case "REPORT_SEGMENT":
			meeting, err := h.roomSvc.SearchMeeting(roomID)
			if err != nil {
				continue
			}

			// Payload contains target_user_id and round_index
			var payload struct {
				TargetUserID string `json:"target_user_id"`
				RoundIndex   int    `json:"round_index"`
			}
			json.Unmarshal(wsMsg.Payload, &payload)

			// Log safety case (for Phase 6 evaluation)
			h.redis.SAdd(h.ctx, fmt.Sprintf("meeting:%s:safety_cases", meeting.ID.String()),
				fmt.Sprintf("%s:%d:%s", payload.TargetUserID, payload.RoundIndex, client.userID.String()))

			log.Printf("Report filed by %s against %s for round %d", client.userID, payload.TargetUserID, payload.RoundIndex)

		case "FORCE_MUTE":
			meeting, err := h.roomSvc.SearchMeeting(roomID)
			if err != nil {
				continue
			}

			isOrganizer, _ := h.roomSvc.IsOrganizer(meeting.ID, client.userID)
			if !isOrganizer {
				continue
			}

			var payload struct {
				TargetUserID string `json:"target_user_id"`
			}
			json.Unmarshal(wsMsg.Payload, &payload)

			h.redis.Publish(h.ctx, "room:"+roomID+":user:"+payload.TargetUserID, `{"type":"FORCE_MUTE"}`)

		case "FORCE_END_MEETING":
			meeting, err := h.roomSvc.SearchMeeting(roomID)
			if err != nil {
				continue
			}
			isOrganizer, _ := h.roomSvc.IsOrganizer(meeting.ID, client.userID)
			if !isOrganizer {
				continue
			}

			h.redis.Publish(h.ctx, "room:"+roomID, `{"type":"FORCE_END_MEETING"}`)
			h.enterGracePeriod(roomID, h.rooms[roomID]) // Trigger grace period

		case "START_INTERVENTION":
			meeting, err := h.roomSvc.SearchMeeting(roomID)
			if err != nil {
				continue
			}
			isOrganizer, _ := h.roomSvc.IsOrganizer(meeting.ID, client.userID)
			if !isOrganizer {
				continue
			}

			h.redis.Set(h.ctx, ball.ActiveKey, client.userID.String(), 0)
			h.broadcastState(roomID)

		case "REQUEST_BALL":
			round, _ := h.ballSvc.GetRound(roomID)
			if round > 1 {
				spokenKey := fmt.Sprintf("room:%s:round:%d:spoken", roomID, round)
				isSpoken, _ := h.redis.SIsMember(h.ctx, spokenKey, client.userID.String()).Result()
				if isSpoken {
					lockKey := fmt.Sprintf("room:%s:user:%s:lock", roomID, client.userID.String())
					_, err := h.redis.Get(h.ctx, lockKey).Result()
					if err == nil {
						log.Printf("User %s is locked from passive queue in round %d", client.userID, round)
						continue
					}

					h.redis.Set(h.ctx, lockKey, 1, 60*time.Second)
				}
			}

			h.ballSvc.RequestBall(roomID, client.userID)
			active, _ := h.ballSvc.GetActiveSpeaker(roomID)
			if active == "" {
				h.passBall(roomID)
			} else {
				h.broadcastState(roomID)
			}
		case "PASS_BALL":
			active, _ := h.ballSvc.GetActiveSpeaker(roomID)
			if active == client.userID.String() {
				h.passBall(roomID)
			}
		default:
			h.redis.Publish(h.ctx, "room:"+roomID, message)
		}
	}
}

func (h *Hub) passBall(roomID string) {
	newSpeaker, _ := h.ballSvc.AssignNextSpeaker(roomID)
	if newSpeaker == "" {
		h.ballSvc.IncrementRound(roomID)
	} else {
		round, _ := h.ballSvc.GetRound(roomID)
		spokenKey := fmt.Sprintf("room:%s:round:%d:spoken", roomID, round)
		h.redis.SAdd(h.ctx, spokenKey, newSpeaker)
		h.redis.Expire(h.ctx, spokenKey, 1*time.Hour)
	}

	h.mutex.Lock()
	room, ok := h.rooms[roomID]
	if ok {
		room.mutex.Lock()
		if room.turnTimer != nil {
			room.turnTimer.Stop()
		}
		if newSpeaker != "" {
			room.turnTimer = time.AfterFunc(ball.TurnLimit, func() {
				log.Printf("Turn timer expired for room %s", roomID)
				h.passBall(roomID)
			})
		}
		room.mutex.Unlock()
	}
	h.mutex.Unlock()

	h.broadcastState(roomID)
}

func (h *Hub) broadcastState(roomID string) {
	active, _ := h.ballSvc.GetActiveSpeaker(roomID)
	queue, _ := h.ballSvc.GetQueue(roomID)

	update := BallStateUpdate{
		Type:   "BALL_STATE",
		Active: active,
		Queue:  queue,
	}

	msg, _ := json.Marshal(update)
	h.redis.Publish(h.ctx, "room:"+roomID, msg)
}

func (h *Hub) subscribeToRoom(roomID string) {
	pubsub := h.redis.Subscribe(h.ctx, "room:"+roomID)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		h.broadcastToRoom(roomID, []byte(msg.Payload))
	}
}

func (h *Hub) broadcastToRoom(roomID string, message []byte) {
	h.mutex.Lock()
	room, ok := h.rooms[roomID]
	h.mutex.Unlock()

	if ok {
		room.mutex.Lock()
		for client := range room.clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(room.clients, client)
			}
		}
		room.mutex.Unlock()
	}
}

func (client *Client) writePump() {
	defer client.conn.Close()
	for {
		message, ok := <-client.send
		if !ok {
			client.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		client.conn.WriteMessage(websocket.TextMessage, message)
	}
}
