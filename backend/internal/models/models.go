package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	PasswordHash     string    `json:"-" db:"password_hash"`
	OrganizerStatus  string    `json:"organizer_status" db:"organizer_status"`
	ProbationCounter int       `json:"probation_counter" db:"probation_counter"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type Meeting struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	ShortID     string     `json:"short_id" db:"short_id"`
	Title       string     `json:"title" db:"title"`
	CohortID    *uuid.UUID `json:"cohort_id,omitempty" db:"cohort_id"`
	CreatorID   uuid.UUID  `json:"creator_id" db:"creator_id"`
	Status      string     `json:"status" db:"status"`
	MeetingType string     `json:"meeting_type" db:"meeting_type"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty" db:"scheduled_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

type MeetingParticipant struct {
	MeetingID   uuid.UUID `json:"meeting_id" db:"meeting_id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Alias       string    `json:"alias" db:"alias"`
	IsOrganizer bool      `json:"is_organizer" db:"is_organizer"`
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
}

type SafetyCase struct {
	ID           uuid.UUID `json:"id" db:"id"`
	MeetingID    uuid.UUID `json:"meeting_id" db:"meeting_id"`
	ReporterID   uuid.UUID `json:"reporter_id" db:"reporter_id"`
	AccusedID    uuid.UUID `json:"accused_id" db:"accused_id"`
	RoundIndex   int       `json:"round_index" db:"round_index"`
	S3EscrowPath string    `json:"s3_escrow_path,omitempty" db:"s3_escrow_path"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
