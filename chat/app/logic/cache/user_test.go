package cache

import (
	"testing"

	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/pkg/redis"
)

func init() {
	conf.Init("../../../conf/config.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func TestUserCache_MultiGetCache(t *testing.T) {
	cache := NewUserCache()
	err := cache.SetUserCache(nil, 1, &model.UserModel{
		Username: "aaa",
	})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	err = cache.SetUserCache(nil, 2, &model.UserModel{
		Username: "bbb",
	})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	data, err := cache.GetUserCache(nil, 1)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("data:%v", data)
	datas, err := cache.MultiGetUserCache(nil, []uint32{1, 2})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("datas:%v", datas)
}
