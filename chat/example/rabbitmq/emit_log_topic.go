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

	err = ch.ExchangeDeclare("logs_topic", "topic",
		true, false, false, false, nil)
	base.FailOnError(err, "failed to declare an exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic", severityFrom(os.Args),
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
	base.FailOnError(err, "failed to publish")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}