package redis

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	// PrefixCheckRepeat check repeat key
	PrefixCheckRepeat = "CHECK_REPEAT"
	// RepeatDefaultTimeout define default timeout
	RepeatDefaultTimeout = 60
)

// CheckRepeat define interface
type CheckRepeat interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Del(ctx context.Context, keys string) int64
}

type checkRepeat struct {
	client *redis.Client
}

// NewCheckRepeat create a check repeat
func NewCheckRepeat(client *redis.Client) CheckRepeat {
	return &checkRepeat{
		client: client,
	}
}

// GetKey 获取key
func getKey(key string) string {
	keyPrefix := "repeat"
	return strings.Join([]string{keyPrefix, PrefixCheckRepeat, key}, ":")
}

func (c *checkRepeat) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	key = getKey(key)
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *checkRepeat) Get(ctx context.Context, key string) (string, error) {
	key = getKey(key)
	return c.client.Get(ctx, key).Result()
}

func (c *checkRepeat) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	key = getKey(key)
	return c.client.SetNX(ctx, key, value, expiration).Result()
}

func (c *checkRepeat) Del(ctx context.Context, key string) int64 {
	key = getKey(key)
	var keys []string
	keys = append(keys, key)
	return c.client.Del(ctx, keys...).Val()
}
