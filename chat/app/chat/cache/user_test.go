package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"chat/pkg/redis"
	"chat/proto/base"
)

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestUserCache_MultiGetCache(t *testing.T) {
	cache := NewUserCache()
	ctx := context.Background()
	err := cache.SetCache(ctx, 1, &base.UserInfo{
		Id:       1,
		Username: "aaa",
	})
	assert.NoError(t, err)
	err = cache.SetCache(ctx, 2, &base.UserInfo{
		Id:       2,
		Username: "bbb",
	})
	assert.NoError(t, err)
	data, err := cache.GetCache(ctx, 1)
	assert.NoError(t, err)
	t.Logf("data:%v", data)
	datas, err := cache.MultiGetCache(ctx, []uint32{1, 2})
	assert.NoError(t, err)
	t.Logf("datas:%v", datas)
}
