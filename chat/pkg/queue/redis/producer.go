package redis

import (
	"chat/pkg/log"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
	"context"
)

//Producer redis消息发布
type Producer struct {
	channel string
}

//NewProducer 创建生成这
func NewProducer(channel string) iqueue.Producer {
	return &Producer{channel: channel}
}

//Publish push消息
func (p *Producer) Publish(ctx context.Context, msg []byte) error {
	return redis.Client.Publish(ctx, p.channel, msg).Err()
}

//MultiPublish 批量push消息
func (p *Producer) MultiPublish(ctx context.Context, msg ...[]byte) (err error) {
	for _, m := range msg {
		err = p.Publish(ctx, m)
		if err != nil {
			log.Warnf("[queue.redis] multi publish err:%v", err)
		}
	}
	return nil
}

//Stop 停止生产者
func (p *Producer) Stop()  {}