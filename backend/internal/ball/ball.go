package ball

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	QueueKey  = "room:%s:ball:queue"
	ActiveKey = "room:%s:ball:active"
)

type BallService struct {
	rdb *redis.Client
	ctx context.Context
}

func NewBallService(rdb *redis.Client) *BallService {
	return &BallService{
		rdb: rdb,
		ctx: context.Background(),
	}
}

// RequestBall adds a user to the room's speaking queue
func (s *BallService) RequestBall(roomID string, userID uuid.UUID) error {
	key := fmt.Sprintf(QueueKey, roomID)
	return s.rdb.RPush(s.ctx, key, userID.String()).Err()
}

// GetActiveSpeaker returns the ID of the user currently holding the ball
func (s *BallService) GetActiveSpeaker(roomID string) (string, error) {
	key := fmt.Sprintf(ActiveKey, roomID)
	return s.rdb.Get(s.ctx, key).Result()
}

// AssignNextSpeaker takes the next user from the queue and makes them active
func (s *BallService) AssignNextSpeaker(roomID string) (string, error) {
	queueKey := fmt.Sprintf(QueueKey, roomID)
	activeKey := fmt.Sprintf(ActiveKey, roomID)

	userID, err := s.rdb.LPop(s.ctx, queueKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // Queue is empty
		}
		return "", err
	}

	err = s.rdb.Set(s.ctx, activeKey, userID, 0).Err()
	if err != nil {
		return "", err
	}

	return userID, nil
}
