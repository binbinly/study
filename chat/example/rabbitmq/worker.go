package main

import (
	"bytes"
	"chat/example/rabbitmq/base"
	"log"
	"time"
)

func main()  {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false, false, false,nil)
	base.FailOnError(err, "failed to declare queue")

	err = ch.Qos(1, 0, false)
	base.FailOnError(err, "failed to set qos")

	msgs, err := ch.Consume(
		q.Name, "", false, false, false, false, nil)
	base.FailOnError(err, "failed to register consume")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}