package ws

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/rifqah/backend/internal/ball"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust for production
	},
}

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	userID uuid.UUID
}

type Room struct {
	clients map[*Client]bool
	mutex   sync.Mutex
}

type Hub struct {
	rooms   map[string]*Room
	mutex   sync.Mutex
	redis   *redis.Client
	ctx     context.Context
	ballSvc *ball.BallService
}

func NewHub(rdb *redis.Client, ballSvc *ball.BallService) *Hub {
	return &Hub{
		rooms:   make(map[string]*Room),
		redis:   rdb,
		ctx:     context.Background(),
		ballSvc: ballSvc,
	}
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
	room.clients[client] = true
	room.mutex.Unlock()
	h.mutex.Unlock()
}

func (h *Hub) unregisterClient(roomID string, client *Client) {
	h.mutex.Lock()
	if room, ok := h.rooms[roomID]; ok {
		room.mutex.Lock()
		if _, ok := room.clients[client]; ok {
			delete(room.clients, client)
			close(client.send)

			// Ghost Ball Protection: Release ball if this client held it
			activeID, _ := h.ballSvc.GetActiveSpeaker(roomID)
			if activeID == client.userID.String() {
				log.Printf("Current speaker %s disconnected, passing the ball", activeID)
				h.ballSvc.AssignNextSpeaker(roomID)
			}
		}
		if len(room.clients) == 0 {
			delete(h.rooms, roomID)
		}
		room.mutex.Unlock()
	}
	h.mutex.Unlock()
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
		
		// Publish to Redis
		err = h.redis.Publish(h.ctx, "room:"+roomID, message).Err()
		if err != nil {
			log.Printf("redis publish error: %v", err)
		}
	}
}

func (client *Client) writePump() {
	defer func() {
		client.conn.Close()
	}()

	for {
		message, ok := <-client.send
		if !ok {
			client.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := client.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		if err := w.Close(); err != nil {
			return
		}
	}
}
