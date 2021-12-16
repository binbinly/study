package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
	"log"
	"time"
)

var (
	topic = "stock.locked"
)

func pub(b broker.Broker)  {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body:   []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func main()  {
	cmd.Init()

	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@192.168.8.76:5672"),
		rabbitmq.ExchangeName("stock.event.exchange"),
	)

	if err := b.Init();err != nil {
		log.Fatalf("Broker init error: %v", err)
	}

	if err := b.Connect(); err != nil {
		log.Fatalf("broker connect error: %v", err)
	}

	pub(b)
}