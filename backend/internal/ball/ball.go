package ball

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	QueueKey  = "room:%s:ball:queue"
	ActiveKey = "room:%s:ball:active"
	RoundKey  = "room:%s:round"
	TurnLimit = 3 * time.Minute
	KeyTTL    = 1 * time.Hour
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

func (s *BallService) GetRound(roomID string) (int, error) {
	key := fmt.Sprintf(RoundKey, roomID)
	val, err := s.rdb.Get(s.ctx, key).Int()
	if err == redis.Nil {
		s.rdb.Set(s.ctx, key, 1, KeyTTL)
		return 1, nil
	}
	// Refresh TTL on access
	s.rdb.Expire(s.ctx, key, KeyTTL)
	return val, err
}

func (s *BallService) IncrementRound(roomID string) (int64, error) {
	key := fmt.Sprintf(RoundKey, roomID)
	val, err := s.rdb.Incr(s.ctx, key).Result()
	s.rdb.Expire(s.ctx, key, KeyTTL)
	return val, err
}

func (s *BallService) RequestBall(roomID string, userID uuid.UUID) error {
	key := fmt.Sprintf(QueueKey, roomID)
	err := s.rdb.RPush(s.ctx, key, userID.String()).Err()
	s.rdb.Expire(s.ctx, key, KeyTTL)
	return err
}

func (s *BallService) GetActiveSpeaker(roomID string) (string, error) {
	key := fmt.Sprintf(ActiveKey, roomID)
	val, err := s.rdb.Get(s.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	s.rdb.Expire(s.ctx, key, KeyTTL)
	return val, err
}

func (s *BallService) GetQueue(roomID string) ([]string, error) {
	key := fmt.Sprintf(QueueKey, roomID)
	return s.rdb.LRange(s.ctx, key, 0, -1).Result()
}

func (s *BallService) AssignNextSpeaker(roomID string) (string, error) {
	queueKey := fmt.Sprintf(QueueKey, roomID)
	activeKey := fmt.Sprintf(ActiveKey, roomID)

	userID, err := s.rdb.LPop(s.ctx, queueKey).Result()
	if err != nil {
		if err == redis.Nil {
			s.rdb.Del(s.ctx, activeKey)
			return "", nil
		}
		return "", err
	}

	err = s.rdb.Set(s.ctx, activeKey, userID, TurnLimit).Err()
	if err != nil {
		return "", err
	}
	// Set expiry for the active key
	s.rdb.Expire(s.ctx, activeKey, KeyTTL)

	return userID, nil
}
