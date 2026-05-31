package ball

import (
	"testing"

	"github.com/rifqah/backend/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestBallService_RoundManagement(t *testing.T) {
	_, rdb := testutils.SetupMockRedis(t)
	ballSvc := NewBallService(rdb)
	roomID := "testroom"

	// Test Initial Round
	round, err := ballSvc.GetRound(roomID)
	assert.NoError(t, err)
	assert.Equal(t, 1, round)

	// Increment Round
	newRound, err := ballSvc.IncrementRound(roomID)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), newRound)

	// Get Round again
	round, err = ballSvc.GetRound(roomID)
	assert.NoError(t, err)
	assert.Equal(t, 2, round)
}
