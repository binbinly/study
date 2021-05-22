package base

import (
	"log"

	"github.com/streadway/amqp"
)

func GetConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.162.170:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}