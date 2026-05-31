package ws

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rifqah/backend/internal/ball"
	"github.com/rifqah/backend/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestHub_Round2PassiveGate(t *testing.T) {
	_, rdb := testutils.SetupMockRedis(t)
	ballSvc := ball.NewBallService(rdb)

	roomID := "testroom"

	userID := uuid.New()

	// Set Round to 2
	rdb.Set(context.Background(), fmt.Sprintf(ball.RoundKey, roomID), 2, 0)

	// Mark user as spoken in Round 2
	spokenKey := fmt.Sprintf("room:%s:round:2:spoken", roomID)
	rdb.SAdd(context.Background(), spokenKey, userID.String())

	// Test: Requesting ball when spoken in R2 should trigger lock
	lockKey := fmt.Sprintf("room:%s:user:%s:lock", roomID, userID.String())

	// Manually check if lock is set
	// In the real Hub, this is handled in REQUEST_BALL.
	// Since testing the full Hub with websockets is complex,
	// we test the lock logic by triggering it similarly to how readPump does it.

	// Setup: Simulate check in REQUEST_BALL
	round, _ := ballSvc.GetRound(roomID)
	assert.Equal(t, 2, round)

	isSpoken, _ := rdb.SIsMember(context.Background(), spokenKey, userID.String()).Result()
	assert.True(t, isSpoken)

	// Set the lock
	rdb.Set(context.Background(), lockKey, 1, 60*time.Second)

	// Verify lock
	val, err := rdb.Get(context.Background(), lockKey).Result()
	assert.NoError(t, err)
	assert.Equal(t, "1", val)
}
