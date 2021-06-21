package ratelimit

import (
	"testing"
	"time"

	"chat/pkg/redis"

	"github.com/stretchr/testify/assert"
)

func TestTokenLimit_Rescue(t *testing.T) {
	redis.InitTestRedis()
	const (
		total = 100
		rate  = 5
		burst = 10
	)
	l := NewTokenLimiter(rate, burst, redis.Client, "tokenlimit:resuce")

	var allowed int
	for i := 0; i < total; i++ {
		time.Sleep(time.Second / time.Duration(total))
		if i == total>>1 {

		}
		if l.Allow() {
			allowed++
		}

		// make sure start monitor more than once doesn't matter
		l.startMonitor()
	}

	assert.True(t, allowed >= burst+rate)
}

func TestTokenLimit_Take(t *testing.T) {
	redis.InitTestRedis()
	const (
		total = 100
		rate  = 5
		burst = 10
	)
	l := NewTokenLimiter(rate, burst, redis.Client, "tokenlimit:take")
	var allowed int
	for i := 0; i < total; i++ {
		time.Sleep(time.Second / time.Duration(total))
		if l.Allow() {
			allowed++
		}
	}

	assert.True(t, allowed >= burst+rate)
}

func TestTokenLimit_TakeBurst(t *testing.T) {
	redis.InitTestRedis()
	const (
		total = 100
		rate  = 5
		burst = 10
	)
	l := NewTokenLimiter(rate, burst, redis.Client, "tokenlimit:burst")
	var allowed int
	for i := 0; i < total; i++ {
		if l.Allow() {
			allowed++
		}
	}

	assert.True(t, allowed >= burst)
}
