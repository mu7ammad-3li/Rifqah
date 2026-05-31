package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"`
	OrganizerStatus  string    `json:"organizer_status"`
	ProbationCounter int       `json:"probation_counter"`
	CreatedAt        time.Time `json:"created_at"`
}

type Meeting struct {
	ID          uuid.UUID  `json:"id"`
	ShortID     string     `json:"short_id"`
	Title       string     `json:"title"`
	CohortID    *uuid.UUID `json:"cohort_id,omitempty"`
	CreatorID   uuid.UUID  `json:"creator_id"`
	Status      string     `json:"status"`
	MeetingType string     `json:"meeting_type"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	EndedAt     *time.Time `json:"ended_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type MeetingParticipant struct {
	MeetingID   uuid.UUID `json:"meeting_id"`
	UserID      uuid.UUID `json:"user_id"`
	Alias       string    `json:"alias"`
	IsOrganizer bool      `json:"is_organizer"`
	JoinedAt    time.Time `json:"joined_at"`
}

type SafetyCase struct {
	ID           uuid.UUID `json:"id"`
	MeetingID    uuid.UUID `json:"meeting_id"`
	ReporterID   uuid.UUID `json:"reporter_id"`
	AccusedID    uuid.UUID `json:"accused_id"`
	RoundIndex   int       `json:"round_index"`
	S3EscrowPath string    `json:"s3_escrow_path,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
