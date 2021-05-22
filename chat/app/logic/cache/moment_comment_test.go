package cache

import (
	"chat/app/logic/conf"
	logger "chat/pkg/log"
	"chat/pkg/redis"
	"context"
	"testing"
)


func init()  {
	conf.Init("../../../../config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
	logger.InitLog(&conf.Conf.Logger)
}

func TestCommentCache_GetCache(t *testing.T) {
	cache := NewCommentCache()
	data, err := cache.MultiGetCache(context.TODO(), []uint32{1,2,3})
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	t.Logf("data:%#v", data)
}
