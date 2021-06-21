package ratelimit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"chat/pkg/redis"
)

func TestPeriodLimit_Take(t *testing.T) {
	testPeriodLimit(t)
}

func TestPeriodLimit_TakeWithAlign(t *testing.T) {
	testPeriodLimit(t, Align())
}

func TestPeriodLimit_RedisUnavailable(t *testing.T) {
	redis.InitTestRedis()
	const (
		seconds = 1
		total   = 100
		quota   = 5
	)

	l := NewPeriodLimit(seconds, quota, redis.Client, "periodlimit1")

	val, err := l.Take("first")
	assert.Nil(t, err)
	assert.Equal(t, 1, val)
}

func testPeriodLimit(t *testing.T, opts ...PeriodOption) {
	redis.InitTestRedis()
	const (
		seconds = 1
		total   = 100
		quota   = 5
	)
	key := fmt.Sprintf("periodlimit:%d", len(opts))
	l := NewPeriodLimit(seconds, quota, redis.Client, key, opts...)
	var allowed, hitQuota, overQuota int
	for i := 0; i < total; i++ {
		val, err := l.Take("first")
		if err != nil {
			t.Error(err)
		}
		switch val {
		case Allowed:
			allowed++
		case HitQuota:
			hitQuota++
		case OverQuota:
			overQuota++
		default:
			t.Error("unknown status")
		}
	}

	assert.Equal(t, quota-1, allowed)
	assert.Equal(t, 1, hitQuota)
	assert.Equal(t, total-quota, overQuota)
}
