package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"mall/pkg/redis"
)

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func Test_redisCache_SetGet(t *testing.T) {

	// 获取redis客户端
	redisClient := redis.Client
	// 实例化redis cache
	cache := NewRedisCache(redisClient, "unit-test", JSONEncoding{}, func() interface{} {
		return new(int64)
	})

	// test set
	type setArgs struct {
		key        string
		value      interface{}
		expiration time.Duration
	}

	setTests := []struct {
		name    string
		cache   Driver
		args    setArgs
		wantErr bool
	}{
		{
			"test redis set",
			cache,
			setArgs{"key-001", "val-001", 60 * time.Second},
			false,
		},
	}

	for _, tt := range setTests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.cache
			if err := c.Set(context.TODO(), tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// test get
	type args struct {
		key string
	}

	tests := []struct {
		name    string
		cache   Driver
		args    args
		wantVal interface{}
		wantErr bool
	}{
		{
			"test redis get",
			cache,
			args{"key-001"},
			"val-001",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.cache
			var gotVal interface{}
			err := c.Get(context.TODO(), tt.args.key, &gotVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotVal, tt.wantVal)
		})
	}
}
