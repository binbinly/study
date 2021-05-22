package main

import (
	"chat/example/rabbitmq/base"
	"log"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout",
		true, false, false, false, nil)
	base.FailOnError(err, "failed to declare exchange")

	q, err := ch.QueueDeclare("", false,
		false, true, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	err = ch.QueueBind(q.Name, "", "logs",
		false, nil)
	base.FailOnError(err, "failed to bind queue")

	msgs, err := ch.Consume(q.Name, "",
		true, false, false, false, nil)
	base.FailOnError(err, "failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}