package redis

import (
	"context"
	"testing"
	"time"
)

func TestInitTestRedis(t *testing.T) {
	InitTestRedis()

	err := Client.Ping(context.Background()).Err()
	if err != nil {
		t.Error("ping redis server err: ", err)
		return
	}
	t.Log("ping redis server pass")
}

func TestRedisSetGet(t *testing.T) {
	InitTestRedis()

	var setGetKey = "test-set"
	var setGetValue = "test-content"
	Client.Set(context.Background(), setGetKey, setGetValue, time.Second*100)

	expectValue := Client.Get(context.Background(), setGetKey).Val()
	if setGetValue != expectValue {
		t.Log("original value: ", setGetValue)
		t.Log("expect value: ", expectValue)
		return
	}

	t.Log("redis set get test pass")
}
