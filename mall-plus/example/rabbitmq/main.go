package main

import (
	"github.com/streadway/amqp"
	"log"
)

const (
	routeKey = "order.create.order"
	exchange = "order.event.exchange"
	queueDelay = "order.delay.queue"
	queue = "order.release.queue"
)

func main()  {
	conn := getConn()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	delayQ, err := ch.QueueDeclare(
		"order.delay.queue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp.Table{
			"x-dead-letter-exchange": exchange,	// 设置死信交换器
			"x-dead-letter-routing-key": "order.release.order",	// 设置死信路由键
			"x-message-ttl": 60000,	//设置过期时间
		},     // arguments
	)
	failOnError(err, "failed to declare a queue")

	releaseQ, err := ch.QueueDeclare(
		"order.release.queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "failed to declare a queue")

	err = ch.ExchangeDeclare(exchange, "topic",
		true, false,
		false, false, nil)
	failOnError(err, "failed to declare an exchange")

	err = ch.QueueBind(delayQ.Name, "order.create.order", exchange,
		false, nil)
	failOnError(err, "failed to bind queue")

	err = ch.QueueBind(releaseQ.Name, "order.release.order", exchange,
		false, nil)
	failOnError(err, "failed to bind queue")

	err = ch.Publish(
		exchange, "order.create.order",
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("hello world"),
		})
	failOnError(err, "failed publish message")

	log.Println("publish success")

	msgs, err := ch.Consume(releaseQ.Name, "", true,
		false, false,
		false, nil)
	failOnError(err, "failed to register consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x]: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func getConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.8.76:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}