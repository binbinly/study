package main

import (
	"chat/example/rabbitmq/base"
	"log"

	"github.com/streadway/amqp"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	base.FailOnError(err, "Failed to declare a queue")

	body := "hello world"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	base.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

