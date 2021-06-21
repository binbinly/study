package main

import (
	logger "chat/pkg/log"
	"chat/pkg/queue/rabbitmq"
	"context"
	"log"
)

func main()  {
	logger.InitLog(logger.NewConfig())
	addr := "guest:guest@192.168.162.170:5672/"

	exchangeName := "test-exchange"
	queueName := "test-bind-to-exchange"

	var message = []byte("Hello World RabbitMQ!")

	config := &rabbitmq.Config{
		Addr:     addr,
		QueueName: queueName,
		Exchange: exchangeName,
		ExchangeType: "fanout",
	}
	go func() {
		producer := rabbitmq.NewProducer(config)
		defer producer.Stop()
		if err := producer.Start(); err != nil {
			log.Printf("start producer err: %s", err.Error())
		}
		if err := producer.Publish(context.Background(), message); err != nil {
			log.Printf("failed publish message: %s", err.Error())
		}
	}()

	// 自定义消息处理函数
	handler := func(body []byte) error {
		log.Println("consumer handler receive msg: ", string(body))
		return nil
	}

	go func() {
		consumer := rabbitmq.NewConsumer(config, handler)
		defer consumer.Stop()
		if err := consumer.Start(); err != nil {
			log.Printf("failed consume: %s", err)
		}
	}()
	select {}
}