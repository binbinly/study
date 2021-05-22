package cache

import (
	"testing"

	"chat/app/logic/conf"
	"chat/pkg/redis"
)

func init()  {
	conf.Init("../../../config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func TestTimelineCache_GetCache(t *testing.T) {
	cache := NewTimelineCache()
	err := cache.SetCache(nil, 1, 2, 10)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	err = cache.SetCache(nil, 3, 4, 20)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	data, err := cache.GetCache(nil, 1, 3)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("data:%v", data)
}