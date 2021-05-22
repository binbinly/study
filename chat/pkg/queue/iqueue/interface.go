package iqueue

import (
	"chat/pkg/queue/nsq"
	"chat/pkg/queue/rabbitmq"
)

type Config struct {
	Plugin   string
	Channel  string
	Nsq      nsq.Config
	Rabbitmq rabbitmq.Config
}

// Producer queue producer
type Producer interface {
	Publish(msg []byte) error
	MultiPublish(msg ...[]byte) error
	Stop()
}

// Consumer queue consumer
type Consumer interface {
	Stop()
}
