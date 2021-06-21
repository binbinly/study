package app

import (
	"context"
	"fmt"
	"time"

	"chat/pkg/crypt"
	"chat/pkg/redis"
)

// CRepeat define struct 重复检测
type CRepeat struct {
	cRepeatClient redis.CheckRepeat
}

// NewCRepeat create a check repeat
func NewCRepeat() *CRepeat {
	return &CRepeat{cRepeatClient: redis.NewCheckRepeat(redis.Client)}
}

// getKey return a check repeat key
func (c *CRepeat) getKey(userID int, check string) string {
	return crypt.Md5ToString(fmt.Sprintf("%d:%s", userID, check))
}

// Set record a repeat value
func (c *CRepeat) Set(ctx context.Context, userID int, check string, value interface{}, expiration time.Duration) error {
	return c.cRepeatClient.Set(ctx, c.getKey(userID, check), value, expiration)
}

// SetNX  set
func (c *CRepeat) SetNX(ctx context.Context, userID int, check string, value interface{}, expiration time.Duration) (bool, error) {
	return c.cRepeatClient.SetNX(ctx, c.getKey(userID, check), value, expiration)
}

// Get get value
func (c *CRepeat) Get(ctx context.Context, userID int, check string) (interface{}, error) {
	return c.cRepeatClient.Get(ctx, c.getKey(userID, check))
}

// Del delete
func (c *CRepeat) Del(ctx context.Context, userID int, check string) int64 {
	return c.cRepeatClient.Del(ctx, c.getKey(userID, check))
}
