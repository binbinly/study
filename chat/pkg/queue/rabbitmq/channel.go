package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func NewChannel(conn *amqp.Connection) (ch *amqp.Channel, err error) {
	ch, err = conn.Channel()
	if err != nil {
		return nil, errors.Wrapf(err, "new channel err")
	}
	return
}

