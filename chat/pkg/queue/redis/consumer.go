package redis

import (
	"context"
	"log"

	logger "chat/pkg/log"
	"chat/pkg/redis"
)

//Consumer redis消息订阅
type Consumer struct {}

//NewConsumer 创建消费者
func NewConsumer(channel string, handler func([]byte) error) *Consumer {
	//参数1 频道名 字符串类型
	ctx := context.TODO()
	sub := redis.Client.Subscribe(ctx, channel)
	_, err := sub.Receive(ctx)
	if err != nil {
		log.Fatalf("init redis subscribe err:%v", err)
	}
	go func() {
		ch := sub.Channel()
		for msg := range ch {
			err = handler([]byte(msg.Payload))
			if err != nil {
				logger.Warnf("[queue.redis] handler err:%v msg:%v", msg.Payload)
			}
		}
	}()
	return &Consumer{}
}

//Stop 停止消费者
func (c *Consumer) Stop()  {}