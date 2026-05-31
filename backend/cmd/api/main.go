package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rifqah/backend/internal/auth"
	"github.com/rifqah/backend/internal/ball"
	"github.com/rifqah/backend/internal/db"
	"github.com/rifqah/backend/internal/media"
	"github.com/rifqah/backend/internal/room"
	"github.com/rifqah/backend/internal/ws"
)

type App struct {
	authService  *auth.AuthService
	roomService  *room.RoomService
	ballService  *ball.BallService
	mediaService *media.MediaService
	hub          *ws.Hub
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file, continuing with environment variables")
	}

	database, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	ballSvc := ball.NewBallService(rdb)

	app := &App{
		authService:  auth.NewAuthService(database),
		roomService:  room.NewRoomService(database),
		ballService:  ballSvc,
		mediaService: media.NewMediaService(),
		hub:          ws.NewHub(rdb, ballSvc),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Post("/register", app.handleRegister)
	r.Post("/login", app.handleLogin)

	r.Group(func(r chi.Router) {
		r.Post("/meetings", app.handleCreateMeeting)
		r.Get("/meetings/{shortID}", app.handleSearchMeeting)
		r.Post("/meetings/{meetingID}/join", app.handleJoinMeeting)
		r.Get("/meetings/{meetingID}/token", app.handleGetToken)

		r.Get("/ws/{roomID}", app.hub.HandleWebSocket)

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func (app *App) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := app.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *App) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := app.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *App) handleCreateMeeting(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string     `json:"title"`
		CohortID    *uuid.UUID `json:"cohort_id"`
		CreatorID   uuid.UUID  `json:"creator_id"`
		MeetingType string     `json:"meeting_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	meeting, err := app.roomService.CreateMeeting(req.Title, req.CohortID, req.CreatorID, req.MeetingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

func (app *App) handleSearchMeeting(w http.ResponseWriter, r *http.Request) {
	shortID := chi.URLParam(r, "shortID")
	meeting, err := app.roomService.SearchMeeting(shortID)
	if err != nil {
		http.Error(w, "Meeting not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(meeting)
}

func (app *App) handleJoinMeeting(w http.ResponseWriter, r *http.Request) {
	meetingIDStr := chi.URLParam(r, "meetingID")
	meetingID, err := uuid.Parse(meetingIDStr)
	if err != nil {
		http.Error(w, "Invalid meeting ID", http.StatusBadRequest)
		return
	}

	var req struct {
		UserID uuid.UUID `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	participant, err := app.roomService.JoinMeeting(meetingID, req.UserID, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(participant)
}

func (app *App) handleGetToken(w http.ResponseWriter, r *http.Request) {
	meetingID := chi.URLParam(r, "meetingID")
	userIDStr := r.URL.Query().Get("userID")
	alias := r.URL.Query().Get("alias")

	if meetingID == "" || userIDStr == "" {
		http.Error(w, "MeetingID and UserID required", http.StatusBadRequest)
		return
	}

	token, err := app.mediaService.GenerateJoinToken(meetingID, userIDStr, alias)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
