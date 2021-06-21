package main

import (
	"chat/app/logic/conf"
	"chat/pkg/redis"
	"context"
	"os"
)

func init()  {
	dir, _ := os.Getwd()
	conf.Init(dir + "/config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func main()  {
	redis.Client.Publish(context.Background(),"message", "aaa")
}