package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"

	"chat/pkg/log"
)

type Consumer struct {
	addr          string
	conn          *amqp.Connection
	channel       *amqp.Channel
	queue         amqp.Queue
	connNotify    chan *amqp.Error
	channelNotify chan *amqp.Error
	exchange      string
	exchangeType  string //exchange类型
	routingKey    string
	queueName     string
	consumerTag   string
	autoDelete    bool                    // 是否自动删除
	handler       func(body []byte) error // 业务自定义消费函数
	quit          chan struct{}
}

func NewConsumer(c *Config, handler func(body []byte) error) *Consumer {
	return &Consumer{
		addr:         c.Addr,
		exchange:     c.Exchange,
		exchangeType: c.ExchangeType,
		routingKey:   c.RoutingKey,
		queueName:    c.QueueName,
		consumerTag:  "consumer",
		autoDelete:   c.AutoDelete,
		handler:      handler,
		quit:         make(chan struct{}),
	}
}

func (c *Consumer) Start() error {
	if err := c.Run(); err != nil {
		return err
	}

	go c.ReConnect()

	return nil
}

func (c *Consumer) Stop() {
	close(c.quit)

	if !c.conn.IsClosed() {
		// 关闭 SubMsg message delivery
		if err := c.channel.Cancel(c.consumerTag, true); err != nil {
			log.Warn("rabbitmq consumer - channel cancel failed: ", err)
		}

		if err := c.conn.Close(); err != nil {
			log.Warn("rabbitmq consumer - connection close failed: ", err)
		}
	}
}

func (c *Consumer) Run() error {
	var err error
	if c.conn, err = OpenConnection(c.addr); err != nil {
		return err
	}

	if c.channel, err = NewChannel(c.conn); err != nil {
		c.conn.Close()
		return err
	}

	// bind queue
	if c.queue, err = c.channel.QueueDeclare(c.queueName, true, c.autoDelete, false, false, nil); err != nil {
		c.channel.Close()
		c.conn.Close()
		return err
	}

	if c.exchange != "" {
		if err = NewExchange(c.channel, c.exchange, c.exchangeType); err != nil {
			c.conn.Channel()
			c.conn.Close()
			return err
		}
		if err = c.channel.QueueBind(c.queue.Name, c.routingKey, c.exchange, false, nil); err != nil {
			c.channel.Close()
			c.conn.Close()
			return err
		}
	}

	var delivery <-chan amqp.Delivery
	// NOTE: autoAck param
	delivery, err = c.channel.Consume(c.queueName, c.consumerTag, true, false, false, false, nil)
	if err != nil {
		c.channel.Close()
		c.conn.Close()
		return err
	}

	go c.Handle(delivery)

	c.connNotify = c.conn.NotifyClose(make(chan *amqp.Error))
	c.channelNotify = c.channel.NotifyClose(make(chan *amqp.Error))

	return nil
}

func (c *Consumer) Handle(delivery <-chan amqp.Delivery) {
	for d := range delivery {
		log.Infof("Consumer received a message: %s in queue: %s", d.Body, c.queueName)
		log.Infof("got %dB delivery: [%v]", len(d.Body), d.DeliveryTag)
		go func(delivery amqp.Delivery) {
			if err := c.handler(delivery.Body); err == nil {
				// NOTE: 假如现在有 10 条消息，它们都是并发处理的，如果第 10 条消息最先处理完毕，
				// 那么前 9 条消息都会被 delivery.Ack(true) 给确认掉。后续 9 条消息处理完毕时，
				// 再执行 delivery.Ack(true)，显然就会导致消息重复确认
				// 报 406 PRECONDITION_FAILED 错误， 所以这里为 false
				delivery.Ack(false)
			} else {
				// 重新入队，否则未确认的消息会持续占用内存
				delivery.Reject(true)
			}
		}(d)
	}
	log.Info("handle: async deliveries channel closed")
}

func (c *Consumer) ReConnect() {
	for {
		select {
		case err := <-c.connNotify:
			if err != nil {
				log.Warn("rabbitmq consumer - connection NotifyClose: ", err)
			}
		case err := <-c.channelNotify:
			if err != nil {
				log.Warn("rabbitmq consumer - channel NotifyClose: ", err)
			}
		case <-c.quit:
			return
		}

		// backstop
		if !c.conn.IsClosed() {
			// 关闭 SubMsg message delivery
			if err := c.channel.Cancel(c.consumerTag, true); err != nil {
				log.Warn("rabbitmq consumer - channel cancel failed: ", err)
			}
			if err := c.conn.Close(); err != nil {
				log.Warn("rabbitmq consumer - conn cancel failed: ", err)
			}
		}

		// IMPORTANT: 必须清空 Notify，否则死连接不会释放
		for err := range c.channelNotify {
			log.Info(err)
		}
		for err := range c.connNotify {
			log.Info(err)
		}

	quit:
		for {
			select {
			case <-c.quit:
				return
			default:
				log.Info("rabbitmq consumer - reconnect")

				if err := c.Run(); err != nil {
					log.Warn("rabbitmq consumer - failCheck:", err)

					// sleep 5s reconnect
					time.Sleep(time.Second * 5)
					continue
				}

				break quit
			}
		}
	}

}
