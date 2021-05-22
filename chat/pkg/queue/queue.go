package queue

import (
	"log"

	"chat/pkg/queue/iqueue"
	"chat/pkg/queue/nsq"
	"chat/pkg/queue/rabbitmq"
	"chat/pkg/queue/redis"
)

const (
	pluginRedis    = "redis"
	pluginNsq      = "nsq"
	pluginRabbitmq = "rabbitmq"
	pluginKafka    = "kafka"
)

//queue 生产者
func NewProducer(c *iqueue.Config) iqueue.Producer {
	switch c.Plugin {
	case pluginRedis:
		return redis.NewProducer(c.Channel)
	case pluginNsq:
		return nsq.NewProducer(&c.Nsq)
	case pluginRabbitmq:
		return rabbitmq.NewProducer(&c.Rabbitmq)
	}
	log.Panicf("not found producer plugin:%v", c.Plugin)
	return nil
}

//queue 消费者
func NewConsumer(c *iqueue.Config, handler func([]byte) error) iqueue.Consumer {
	switch c.Plugin {
	case pluginRedis:
		return redis.NewConsumer(c.Channel, handler)
	case pluginNsq:
		return nsq.NewConsumer(&c.Nsq, handler)
	case pluginRabbitmq:
		return rabbitmq.NewConsumer(&c.Rabbitmq, handler)
	}
	log.Fatalf("not found consumer plugin:%v", c.Plugin)
	return nil
}
