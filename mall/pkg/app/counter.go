package app

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

//计数器
const (
	// PrefixCounter counter key
	PrefixCounter = "counter"
	// DefaultStep default step key
	DefaultStep = 1
	// DefaultExpirationTime 默认生存时间
	DefaultExpirationTime = 600 * time.Second
)

// Counter define struct
type Counter struct {
	client *redis.Client
}

// NewCounter create a counter
func NewCounter(rdb *redis.Client) *Counter {
	return &Counter{
		client: rdb,
	}
}

// GetKey 获取key
func (c *Counter) GetKey(key string) string {
	return strings.Join([]string{PrefixCounter, key}, ":")
}

// SetCounter set counter
func (c *Counter) SetCounter(ctx context.Context, idStr string, expiration time.Duration) (int64, error) {
	key := c.GetKey(idStr)
	ret, err := c.client.IncrBy(ctx, key, DefaultStep).Result()
	if err != nil {
		return 0, err
	}
	_, err = c.client.Expire(ctx, key, expiration).Result()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

// GetCounter get total count
func (c *Counter) GetCounter(ctx context.Context, idStr string) (int64, error) {
	return c.client.Get(ctx, c.GetKey(idStr)).Int64()
}

// DelCounter del count
func (c *Counter) DelCounter(ctx context.Context, idStr string) int64 {
	keys := []string{c.GetKey(idStr)}
	return c.client.Del(ctx, keys...).Val()
}
