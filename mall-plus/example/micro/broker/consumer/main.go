package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"log"

	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
)

var (
	topic = "go.micro.topic.foo"
)

func sharedSub(b broker.Broker) {
	_, err := b.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	//brkrSub := broker.NewSubscribeOptions(
	//	broker.Queue("queue.default"),
	//	broker.DisableAutoAck(),
	//	rabbitmq.DurableQueue(),
	//)

	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@192.168.8.76:5672"),
	)

	if err := b.Init(); err != nil {
		log.Fatalf("broker init error:%v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("broker connect error: %v", err)
	}

	sharedSub(b)
	select {}
}
