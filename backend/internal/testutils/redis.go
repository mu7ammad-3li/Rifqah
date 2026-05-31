package testutils

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// SetupMockRedis creates a miniredis instance and returns a redis client connected to it.
func SetupMockRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	return s, rdb
}
