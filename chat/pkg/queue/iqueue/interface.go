package iqueue

import (
	"context"

	"chat/pkg/queue/nsq"
	"chat/pkg/queue/rabbitmq"
)

//Config 队列配置
type Config struct {
	Plugin   string
	Channel  string
	Nsq      nsq.Config
	Rabbitmq rabbitmq.Config
}

// Producer queue producer
type Producer interface {
	Publish(ctx context.Context, msg []byte) error
	MultiPublish(ctx context.Context, msg ...[]byte) error
	Stop()
}

// Consumer queue consumer
type Consumer interface {
	Stop()
}
