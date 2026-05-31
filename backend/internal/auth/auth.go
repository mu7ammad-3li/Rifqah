package auth

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rifqah/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPass  = errors.New("invalid password")
)

type AuthService struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	query := "INSERT INTO users (id, username, email, password_hash, created_at) VALUES (:id, :username, :email, :password_hash, :created_at)"
	
	_, err = s.db.NamedExec(query, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	err := s.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidPass
	}

	return &user, nil
}

var colors = []string{"Blue", "Red", "Green", "Yellow", "Purple", "Orange", "Silver", "Golden", "White", "Black"}
var nouns = []string{"River", "Mountain", "Forest", "Eagle", "Lion", "Wolf", "Star", "Moon", "Sun", "Cloud"}

func GenerateAlias() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return colors[r.Intn(len(colors))] + " " + nouns[r.Intn(len(nouns))]
}
