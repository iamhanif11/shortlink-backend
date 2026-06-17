package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func SaveToBlacklist(ctx context.Context, rc *redis.Client, token string, expiration time.Duration) error {
	rkey := "blacklist:" + token
	return rc.Set(ctx, rkey, "1", expiration).Err()
}

func IsTokenBlacklisted(ctx context.Context, rc *redis.Client, token string) bool {
	rkey := "blacklist:" + token
	val, err := rc.Get(ctx, rkey).Result()
	return err == nil && val == "1"
}
