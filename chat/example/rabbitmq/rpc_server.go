package main

import (
	"chat/example/rabbitmq/base"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("rpc_queue",
		false, false,
		false, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", false,
		false, false, false, nil)
	base.FailOnError(err, "failed to register consume")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			base.FailOnError(err, "failed to convert to integer")

			log.Printf(" [.] fib(%d)", n)

			response := fib(n)

			err = ch.Publish(
				"", d.ReplyTo,
				false, false,
				amqp.Publishing{
					ContentType: "text/plain",
					CorrelationId: d.CorrelationId,
					Body: []byte(strconv.Itoa(response)),
				})
			base.FailOnError(err, "failed to publish")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}