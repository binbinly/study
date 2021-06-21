package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/pkg/redis"
)

func TestMain(m *testing.M) {
	conf.Init("../../../config/logic.yaml")
	redis.Init(&conf.Conf.Redis)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestUserCache_MultiGetCache(t *testing.T) {
	cache := NewUserCache()
	ctx := context.Background()
	err := cache.SetUserCache(ctx, 1, &model.UserModel{
		Username: "aaa",
	})
	assert.NoError(t, err)
	err = cache.SetUserCache(ctx, 2, &model.UserModel{
		Username: "bbb",
	})
	assert.NoError(t, err)
	data, err := cache.GetUserCache(ctx, 1)
	assert.NoError(t, err)
	t.Logf("data:%v", data)
	datas, err := cache.MultiGetUserCache(ctx, []uint32{1, 2})
	assert.NoError(t, err)
	t.Logf("datas:%v", datas)
}
