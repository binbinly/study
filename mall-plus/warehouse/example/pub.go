package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.8.76:5672/")
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	err = ch.Publish(
		"stock.event.exchange", "stock.locked",
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("hello world"),
		})
}