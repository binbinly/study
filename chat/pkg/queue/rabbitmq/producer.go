package rabbitmq

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"

	"chat/pkg/log"
)

const (
	//ExchangeFanout 交换器常用三种模式，
	ExchangeFanout = "fanout"
	//ExchangeDirect direct
	ExchangeDirect = "direct"
	//ExchangeTopic topic
	ExchangeTopic  = "topic"
)

//Producer 生产者
type Producer struct {
	addr          string
	conn          *amqp.Connection
	channel       *amqp.Channel
	queue         amqp.Queue
	queueName     string
	routingKey    string //路由键
	exchange      string //交换器,默认 Work Queues、循环分发、消息确认、持久化、公平分发
	exchangeType  string //exchange类型
	connNotify    chan *amqp.Error
	channelNotify chan *amqp.Error
	quit          chan struct{}
}

//NewProducer 创建生成者
func NewProducer(c *Config) *Producer {
	p := &Producer{
		addr:         c.Addr,
		queueName: c.QueueName,
		exchange:     c.Exchange,
		routingKey:   c.RoutingKey,
		exchangeType: c.ExchangeType,
		quit:         make(chan struct{}),
	}

	return p
}

//Start 开启生产者
func (p *Producer) Start() error {
	if err := p.Run(); err != nil {
		return err
	}
	log.Info("start rabbitmq producer success")
	go p.ReConnect()

	return nil
}

//Stop 体质生产者
func (p *Producer) Stop() {
	close(p.quit)

	if !p.conn.IsClosed() {
		if err := p.conn.Close(); err != nil {
			log.Warnf("rabbitmq producer - connection close failed: %v", err)
		}
	}
}

//Run 运行
func (p *Producer) Run() error {
	var err error
	if p.conn, err = OpenConnection(p.addr); err != nil {
		return err
	}

	if p.channel, err = NewChannel(p.conn); err != nil {
		p.conn.Close()
		return err
	}

	if p.exchange == "" { //默认工作队列模式
		if p.queue, err = NewQueue(p.channel, p.queueName); err != nil {
			p.conn.Channel()
			p.conn.Close()
			return err
		}
	} else {
		if err = NewExchange(p.channel, p.exchange, p.exchangeType); err != nil {
			p.conn.Channel()
			p.conn.Close()
			return err
		}
	}

	p.connNotify = p.conn.NotifyClose(make(chan *amqp.Error))
	p.channelNotify = p.channel.NotifyClose(make(chan *amqp.Error))

	return err
}

//ReConnect 断线重连必读文章 https://ms2008.github.io/2019/06/16/golang-rabbitmq/
func (p *Producer) ReConnect() {
	for {
		select {
		case err := <-p.connNotify:
			if err != nil {
				log.Warnf("rabbitmq producer - connection NotifyClose: %v", err)
			}
		case err := <-p.channelNotify:
			if err != nil {
				log.Warnf("rabbitmq producer - channel NotifyClose: %v", err)
			}
		case <-p.quit:
			return
		}

		// backstop
		if !p.conn.IsClosed() {
			if err := p.conn.Close(); err != nil {
				log.Warnf("rabbitmq producer - connection close failed: %v", err)
			}
		}

		// IMPORTANT: 必须清空 Notify，否则死连接不会释放
		for err := range p.channelNotify {
			log.Info(err)
		}
		for err := range p.connNotify {
			log.Info(err)
		}

	quit:
		for {
			select {
			case <-p.quit:
				return
			default:
				log.Infof("rabbitmq consumer - reconnect")

				if err := p.Run(); err != nil {
					log.Warnf("rabbitmq producer - failCheck: %v", err)

					// sleep 5s reconnect
					time.Sleep(time.Second * 5)
					continue
				}

				break quit
			}
		}
	}
}

//Publish push消息
func (p *Producer) Publish(ctx context.Context, msg []byte) error {
	if p.exchange == "" {
		p.routingKey = p.queue.Name
	}
	return p.channel.Publish(
		p.exchange,   // exchange
		p.routingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			MessageId:    uuid.New().String(),
			Body:         msg,
		})
}

//MultiPublish 批量push消息
func (p *Producer) MultiPublish(ctx context.Context, msg ...[]byte) (err error) {
	for _, m := range msg {
		err = p.Publish(ctx, m)
		if err != nil {
			log.Warnf("[queue.redis] multi publish err:%v", err)
		}
	}
	return nil
}
