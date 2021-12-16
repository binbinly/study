package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//OpenConnection 建立链接
func OpenConnection(addr string) (conn *amqp.Connection, err error) {
	conn, err = amqp.Dial(addr)
	if err != nil {
		return nil, errors.Wrapf(err, "RabbitMQ Open Connection")
	}
	return
}