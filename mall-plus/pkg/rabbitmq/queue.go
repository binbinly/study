package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// NewQueue .
func NewQueue(channel *amqp.Channel, opts ...Option) (q amqp.Queue, err error) {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	q, err = channel.QueueDeclare(
		options.name,
		options.durable,
		options.autoDelete,
		options.exclusive,
		options.noWait,
		options.args,
	)
	if err != nil {
		return q, errors.Wrapf(err, "RabbitMQ New Queue")
	}
	return
}

