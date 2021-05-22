package nsq

import (
	"log"

	"github.com/nsqio/go-nsq"

	logger "chat/pkg/log"
)

type Consumer struct {
	Con     *nsq.Consumer
	handler func(body []byte) error
}

// NewConsumer new一个nsq消费者
func NewConsumer(c *Config, handler func([]byte) error) *Consumer {
	consumer, err := nsq.NewConsumer(c.Topic, c.Channel, setting(c))
	if err != nil {
		log.Fatalf("[newConsumer] new consumer err:%v,topic:%v,channel:%v", err, c.Topic, c.Channel)
	}
	con := &Consumer{handler: handler, Con: consumer}
	consumer.AddHandler(con)
	err = consumer.ConnectToNSQLookupds(c.ConsumerHost)
	if err != nil {
		log.Fatalf("[NewMessageConsumer] err:%v", err)
	}
	log.Printf("[NewMessageConsumer] success topic:%s, channel:%s", c.Topic, c.Channel)
	return con
}

// HandleMessage 处理消息
func (c *Consumer) HandleMessage(m *nsq.Message) (err error) {
	// 重试次数判断 nsq推送过来的消息里，有个attempts字段，代表着尝试的次数，一开始是1，每次客户端给nsq会REQ响应后，nsq再次推送过来的消息，attempts都会加1，
	if m.Attempts > 3 {
		logger.Errorf("[HandleMessage] nsq err :%v, attempts body:%v", err, string(m.Body))
		return nil
	}
	logger.Infof("[HandleMessage] handle message:%v", string(m.Body))
	err = c.handler(m.Body)
	if err != nil {
		logger.Errorf("[HandleMessage] nsq err :%v, body:%v", err, string(m.Body))
	}
	return err
}

func (c *Consumer) Stop() {
	c.Con.Stop()
}
