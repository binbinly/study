package main

import (
	"chat/example/rabbitmq/base"
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout",
		true, false, false,
		false, nil)
	base.FailOnError(err, "failed to declare an exchange")

	err = ch.Publish("logs", "",
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("hello world"),
		})
	base.FailOnError(err, "failed to publish")

	log.Print("publish success")
}