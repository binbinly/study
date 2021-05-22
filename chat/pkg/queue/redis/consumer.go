package redis

import (
	"log"

	logger "chat/pkg/log"
	"chat/pkg/redis"
)

//Consumer redis消息订阅
type Consumer struct {}

func NewConsumer(channel string, handler func([]byte) error) *Consumer {
	//参数1 频道名 字符串类型
	sub := redis.Client.Subscribe(channel)
	_, err := sub.Receive()
	if err != nil {
		log.Fatalf("init redis subscribe err:%v", err)
	}
	go func() {
		for msg := range sub.Channel() {
			err = handler([]byte(msg.Payload))
			if err != nil {
				logger.Warnf("[queue.redis] handler err:%v msg:%v", msg.Payload)
			}
		}
	}()
	return &Consumer{}
}

func (c *Consumer) Stop()  {}