package room

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rifqah/backend/internal/auth"
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

func (s *RoomService) JoinMeeting(meetingID uuid.UUID, userID uuid.UUID, isOrganizer bool) (*models.MeetingParticipant, error) {
	participant := &models.MeetingParticipant{
		MeetingID:   meetingID,
		UserID:      userID,
		Alias:       auth.GenerateAlias(),
		IsOrganizer: isOrganizer,
		JoinedAt:    time.Now(),
	}

	err := s.db.Get(&participant.Alias, "INSERT INTO meeting_participants (meeting_id, user_id, alias, is_organizer, joined_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (meeting_id, user_id) DO UPDATE SET joined_at = $5 RETURNING alias", 
		participant.MeetingID, participant.UserID, participant.Alias, participant.IsOrganizer, participant.JoinedAt)
	
	if err != nil {
		return nil, err
	}

	return participant, nil
}

func generateShortID() (string, error) {
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
