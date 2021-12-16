package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//NewExchange 实例化交换器
func NewExchange(ch *amqp.Channel, exchangeType string, opts ...Option) (err error) {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	err = ch.ExchangeDeclare(
		options.name,
		exchangeType,
		options.durable,
		options.autoDelete,
		options.exclusive,
		options.noWait,
		options.args,
	)
	if err != nil {
		return errors.Wrapf(err, "RabbitMQ New Exchnage")
	}
	return
}