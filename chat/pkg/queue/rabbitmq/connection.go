package rabbitmq

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/streadway/amqp"
)

//OpenConnection 建立链接
func OpenConnection(addr string) (conn *amqp.Connection, err error) {
	uri := fmt.Sprintf("amqp://%s", addr)

	conn, err = amqp.Dial(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "open connection err")
	}
	return
}
