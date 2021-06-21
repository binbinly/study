package kafka

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

//Producer 生产者结构
type Producer struct {
	asyncProducer sarama.AsyncProducer
	topic         string
	enqueued      int
}

//NewProducer 创建生产者
func NewProducer(config *sarama.Config, logger *log.Logger, topic string, brokers []string) *Producer {
	sarama.Logger = logger

	// Start a new async producer
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	log.Println("Kafka AsyncProducer up and running!")

	return &Producer{
		asyncProducer: producer,
		topic:         topic,
	}
}

//Publish push消息
func (p *Producer) Publish(message string) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		time.Sleep(5 * time.Second)
		msg := &sarama.ProducerMessage{Topic: p.topic, Value: sarama.StringEncoder(message)}

		select {
		case p.asyncProducer.Input() <- msg:
			p.enqueued++
			log.Printf("New message publish:  %s", msg.Value)
		case <-signals:
			p.asyncProducer.AsyncClose() // Trigger a shutdown of the producer.
			log.Printf("Kafka AsyncProducer finished with %d messages produced.", p.enqueued)
			return
		}
	}

}
