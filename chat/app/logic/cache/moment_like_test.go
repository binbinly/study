package cache

import (
	"chat/app/logic/conf"
	"chat/pkg/redis"
	"testing"
)

func init()  {
	conf.Init("../../../../config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func TestLikeCache_MultiGetCache(t *testing.T) {
	cache := NewLikeCache()
	err := cache.SetCache(nil, 1, []uint32{1,2,3})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	err = cache.SetCache(nil, 2, []uint32{4,5,6})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	data, err := cache.GetCache(nil, 1)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("data:%v:%v", data, len(*data))
	datas, err := cache.MultiGetCache(nil, []uint32{1,2})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("datas:%#v", datas)
}