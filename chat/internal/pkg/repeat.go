package pkg

import (
	"fmt"
	"time"

	"chat/pkg/crypt"
	"chat/pkg/redis"
)

// CRepeat define struct
type CRepeat struct {
	cRepeatClient redis.CheckRepeat
}

// NewCRepeat create a check repeat
func NewCRepeat() *CRepeat {
	return &CRepeat{cRepeatClient: redis.NewCheckRepeat(redis.Client)}
}

// getKey return a check repeat key
func (c *CRepeat) getKey(userId int, check string) string {
	return crypt.Md5ToString(fmt.Sprintf("%d:%s", userId, check))
}

// Set record a repeat value
func (c *CRepeat) Set(userId int, check string, value interface{}, expiration time.Duration) error {
	return c.cRepeatClient.Set(c.getKey(userId, check), value, expiration)
}

// SetNX  set
func (c *CRepeat) SetNX(userId int, check string, value interface{}, expiration time.Duration) (bool, error) {
	return c.cRepeatClient.SetNX(c.getKey(userId, check), value, expiration)
}

// Get get value
func (c *CRepeat) Get(userId int, check string) (interface{}, error) {
	return c.cRepeatClient.Get(c.getKey(userId, check))
}

// Del delete
func (c *CRepeat) Del(userId int, check string) int64 {
	return c.cRepeatClient.Del(c.getKey(userId, check))
}
