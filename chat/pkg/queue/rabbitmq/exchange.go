package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//NewExchange 实例化交换器
func NewExchange(ch *amqp.Channel, name, t string) (err error) {
	err = ch.ExchangeDeclare(
		name,
		t,
		true,	//是否持久化
		false, //自动删除，在最后一个consumer断开连接后，删除
		false, //独占，只能被一个consumer 的conn使用
		false,
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "new exchnage err")
	}
	return
}