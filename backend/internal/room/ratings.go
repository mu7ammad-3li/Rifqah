package room

import (
	"github.com/google/uuid"
)

// SubmitRating adds an anonymized rating for an organizer
func (s *RoomService) SubmitRating(organizerID uuid.UUID, meetingID uuid.UUID, rating int) error {
	query := "INSERT INTO organizer_ratings_tally (organizer_id, meeting_id, rating) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, organizerID, meetingID, rating)
	return err
}
