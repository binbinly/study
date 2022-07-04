package lock

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	// RedisLockKey redis lock key
	RedisLockKey = "app:redis:lock:%s"
)

// Lock define common func
type Lock interface {
	Lock(ctx context.Context, timeout time.Duration) (bool, error)
	Unlock(ctx context.Context) (bool, error)
}

// genToken 生成token
func genToken() string {
	u, _ := uuid.NewRandom()
	return u.String()
}