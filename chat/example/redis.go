package main

import (
	"fmt"

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
	Subscribe()
}

func Subscribe(){
	//参数1 频道名 字符串类型
	pubsub := redis.Client.Subscribe("message")
	_, err := pubsub.Receive()
	if err != nil {
		return
	}
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println( msg.Channel, msg.Payload)
	}
}