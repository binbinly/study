package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//NewQueue 创建队列
func NewQueue(channel *amqp.Channel, name string) (q amqp.Queue, err error) {
	q, err = channel.QueueDeclare(
		name,
		true,	//是否持久化
		false, //自动删除，在最后一个consumer断开连接后，删除
		false, //独占，只能被一个consumer 的conn使用
		false,
		nil,
	)
	if err != nil {
		return q, errors.Wrapf(err, "new queue err")
	}
	return
}
