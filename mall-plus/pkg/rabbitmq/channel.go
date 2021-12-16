package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//NewChannel 创建信道
func NewChannel(conn *amqp.Connection) (ch *amqp.Channel, err error) {
	ch, err = conn.Channel()
	if err != nil {
		return nil, errors.Wrapf(err, "RabbitMQ New Channel")
	}
	return
}