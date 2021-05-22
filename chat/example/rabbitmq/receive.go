package main

import (
	"chat/example/rabbitmq/base"
	"log"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil)
	base.FailOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "",
		true, false, false, false, nil)
	base.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}