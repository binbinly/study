package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

//Consumer 消费者结构
type Consumer struct {
	group   sarama.ConsumerGroup
	topics  []string
	groupID string
}

//NewConsumer 创建消费者
func NewConsumer(config *sarama.Config, logger *log.Logger, topic string, groupID string, brokers []string) *Consumer {
	// Init config, specify appropriate versio
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	sarama.Logger = logger
	config.Version = sarama.V2_0_0_0 // V2_4_0_0

	// Start with a client
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		panic(err)
	}

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		panic(err)
	}

	log.Println("Consumer up and running")

	return &Consumer{
		group:   group,
		topics:  []string{topic},
		groupID: groupID,
	}
}

//Consume 启动消费者
func (c Consumer) Consume() {
	// Track errors
	go func() {
		for err := range c.group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := ConsumerGroupHandler{}

		err := c.group.Consume(ctx, c.topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

//Stop 停止消费者
func (c Consumer) Stop() {
	c.group.Close()
}
