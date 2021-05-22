package main

import (
	"chat/example/rabbitmq/base"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn := base.GetConn()
	defer conn.Close()

	ch, err := conn.Channel()
	base.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("", false,
		false, true, false, nil)
	base.FailOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "",
		true, false, false, false, nil)
	base.FailOnError(err, "failed to register a consume")

	corrId := randomString(32)

	err = ch.Publish(
		"", "rpc_queue",
		false, false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	base.FailOnError(err, "failed to publish a publish")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			base.FailOnError(err, "Failed to convert body to integer")
			break
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := 10

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	base.FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}
