package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitTestRedis(t *testing.T) {
	err := Client.Ping(context.Background()).Err()
	if err != nil {
		t.Error("ping redis server err: ", err)
		return
	}
	t.Log("ping redis server pass")
}

func TestRedisSetGet(t *testing.T) {
	ctx := context.Background()
	var setGetKey = "test-set"
	var setGetValue = "test-content"
	Client.Set(ctx, setGetKey, setGetValue, time.Second*100)

	expectValue := Client.Get(ctx, setGetKey).Val()
	assert.Equal(t, setGetValue, expectValue)
}
