package main

import (
	"context"
	"fmt"
	"time"

	"chat/app/logic/conf"
	"chat/pkg/redis"
	"os"
)

func init()  {
	dir, _ := os.Getwd()
	conf.Init(dir + "/config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func main()  {
	//Subscribe()
	Pipe()
}

func Pipe()  {
	pipe := redis.Client.Pipeline()
	c := context.Background()
	key := "test"
	pipe.Incr(c, key)
	pipe.Expire(c, key, time.Minute)
	_, err := pipe.Exec(c)
	if err != nil {
		fmt.Println("err", err)
	}

}

func Subscribe(){
	//参数1 频道名 字符串类型
	sub := redis.Client.Subscribe(context.Background(), "message")
	_, err := sub.Receive(context.Background())
	if err != nil {
		return
	}
	ch := sub.Channel()
	for msg := range ch {
		fmt.Println( msg.Channel, msg.Payload)
	}
}