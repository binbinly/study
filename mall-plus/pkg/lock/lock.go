package lock

import (
	"context"
	"time"
)

const (
	// RedisLockKey redis lock key
	RedisLockKey = "eagle:lock:%s"
	// EtcdLockKey etcd lock key
	EtcdLockKey = "/eagle/lock/%s"
)

// Lock 分布式锁
type Lock interface {
	Lock(ctx context.Context, timeout time.Duration) (bool, error)
	Unlock(ctx context.Context) (bool, error)
}
