package main

import (
	"chat/example/rabbitmq/base"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("task_queue",
		true,
		false,
		false, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	body := bodyForm(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //持久化消息
			ContentType: "text/plain",
			Body: []byte(body),
		})
	base.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}