package redis

import (
	"chat/pkg/log"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
)

//Producer redis消息发布
type Producer struct {
	channel string
}

func NewProducer(channel string) iqueue.Producer {
	return &Producer{channel: channel}
}

func (p *Producer) Publish(msg []byte) error {
	return redis.Client.Publish(p.channel, msg).Err()
}

func (p *Producer) MultiPublish(msg ...[]byte) (err error) {
	for _, m := range msg {
		err = p.Publish(m)
		if err != nil {
			log.Warnf("[queue.redis] multi publish err:%v", err)
		}
	}
	return nil
}

func (p *Producer) Stop()  {}