package room

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rifqah/backend/internal/models"
)

type RoomService struct {
	db *sqlx.DB
}

func NewRoomService(db *sqlx.DB) *RoomService {
	return &RoomService{db: db}
}

func (s *RoomService) CreateMeeting(title string, cohortID *uuid.UUID, creatorID uuid.UUID, meetingType string) (*models.Meeting, error) {
	shortID, err := generateShortID()
	if err != nil {
		return nil, err
	}

	meeting := &models.Meeting{
		ID:          uuid.New(),
		ShortID:     shortID,
		Title:       title,
		CohortID:    cohortID,
		CreatorID:   creatorID,
		Status:      "scheduled",
		MeetingType: meetingType,
		CreatedAt:   time.Now(),
	}

	query := "INSERT INTO meetings (id, short_id, title, cohort_id, creator_id, status, meeting_type, created_at) VALUES (:id, :short_id, :title, :cohort_id, :creator_id, :status, :meeting_type, :created_at)"

	_, err = s.db.NamedExec(query, meeting)
	if err != nil {
		return nil, err
	}

	// Creator automatically joins as organizer
	_, err = s.JoinMeeting(meeting.ID, creatorID, true)
	if err != nil {
		return nil, err
	}

	return meeting, nil
}

func (s *RoomService) SearchMeeting(shortID string) (*models.Meeting, error) {
	var meeting models.Meeting
	err := s.db.Get(&meeting, "SELECT * FROM meetings WHERE short_id = $1", shortID)
	if err != nil {
		return nil, err
	}
	return &meeting, nil
}

// ... existing imports

func (s *RoomService) IsOrganizer(meetingID uuid.UUID, userID uuid.UUID) (bool, error) {
	var isOrganizer bool
	err := s.db.Get(&isOrganizer, "SELECT is_organizer FROM meeting_participants WHERE meeting_id = $1 AND user_id = $2", meetingID, userID)
	if err != nil {
		return false, err
	}
	return isOrganizer, nil
}

func generateShortID() (string, error) {
	// ...
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numbers = "0123456789"

	ret := make([]byte, 9)
	for i := 0; i < 4; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}
	ret[4] = '-'
	for i := 5; i < 9; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		ret[i] = numbers[num.Int64()]
	}

	return string(ret), nil
}
