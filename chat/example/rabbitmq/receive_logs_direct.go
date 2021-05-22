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

	err = ch.ExchangeDeclare("logs_direct", "direct",
		true, false, false, false,
		nil)
	base.FailOnError(err, "failed to declare an exchange")

	q, err := ch.QueueDeclare("", false, false,
		true, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	for _, s := range []string{"warn", "info"} {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(q.Name, s, "logs_direct",
			false, nil)
		base.FailOnError(err, "failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false,
		false, nil)
	base.FailOnError(err, "failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}