package pkg

import (
	"strings"
	"time"

	"github.com/go-redis/redis"

	redis2 "chat/pkg/redis"
)

const (
	// PrefixCounter counter key
	PrefixCounter = "counter"
	// DefaultStep default step key
	DefaultStep           = 1
	// DefaultExpirationTime 默认生存时间
	DefaultExpirationTime = 600 * time.Second
)

// Counter define struct
type Counter struct {
	client *redis.Client
}

// NewCounter create a counter
func NewCounter() *Counter  {
	return &Counter{
		client: redis2.Client,
	}
}

// GetKey 获取key
func (c *Counter) GetKey(key string) string {
	return strings.Join([]string{PrefixCounter, key}, ":")
}

// SetCounter set counter
func (c *Counter) SetCounter(idStr string, expiration time.Duration) (int64, error) {
	key := c.GetKey(idStr)
	ret, err := c.client.IncrBy(key, DefaultStep).Result()
	if err != nil {
		return 0, err
	}
	_, err = c.client.Expire(key, expiration).Result()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

// GetCounter get total count
func (c *Counter) GetCounter(idStr string) (int64, error) {
	return c.client.Get(c.GetKey(idStr)).Int64()
}

// DelCounter del count
func (c *Counter) DelCounter(idStr string) int64 {
	keys := []string{c.GetKey(idStr)}
	return c.client.Del(keys...).Val()
}