package main

import (
	"chat/example/rabbitmq/base"
	"log"
	"os"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs_topic", "topic",
		true, false,
		false, false, nil)
	base.FailOnError(err, "failed to declare an exchange")

	q, err := ch.QueueDeclare("", false, false,
		true, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_topic", s)
		err = ch.QueueBind(q.Name, s, "logs_topic",
			false, nil)
		base.FailOnError(err, "failed to bind queue")
	}

	msgs, err := ch.Consume(q.Name, "", true,
		false, false,
		false, nil)
	base.FailOnError(err, "failed to register consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x]: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}