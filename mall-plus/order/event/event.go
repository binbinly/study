package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"go-micro.dev/v4/logger"

	"common/constvar"
	"common/message"
	pb "common/proto/order"
	"order/service"
	"pkg/rabbitmq"
)

const (
	queue      = "order.release.queue"
	delayQueue = "order.delay.queue"
)

//Event All methods of Event will be executed when a message is received
type Event struct {
	srv service.IService
}

//New 实例化
func New(srv service.IService) *Event {
	return &Event{srv: srv}
}

//Init 初始化消费者
func (e *Event) Init(addr string) {
	consumer := rabbitmq.NewConsumer(addr, queue, func(ch *amqp.Channel) error {
		//topic交换机
		err := rabbitmq.NewExchange(ch, amqp.ExchangeTopic,
			rabbitmq.WithName(constvar.ExchangeOrder),
			rabbitmq.WithDurable())
		if err != nil {
			log.Fatal(err)
		}
		//普通队列
		_, err = rabbitmq.NewQueue(ch, rabbitmq.WithName(queue),
			rabbitmq.WithDurable())
		if err != nil {
			log.Fatal(err)
		}
		//延迟队列
		_, err = rabbitmq.NewQueue(ch, rabbitmq.WithName(delayQueue),
			rabbitmq.WithDurable(),
			rabbitmq.WithArgs(amqp.Table{
				"x-dead-letter-exchange":    constvar.ExchangeOrder,   // 设置死信交换器
				"x-dead-letter-routing-key": constvar.KeyOrderRelease, // 设置死信路由键
				"x-message-ttl":             1800000,                  //设置过期时间
			}))
		if err != nil {
			log.Fatal(err)
		}
		//绑定队列
		if err = ch.QueueBind(queue, constvar.KeyOrderRelease, constvar.ExchangeOrder, false, nil); err != nil {
			log.Fatalf("RabbitMq Bind Queue err:%v", err)
		}
		if err = ch.QueueBind(delayQueue, constvar.KeyOrderCreate, constvar.ExchangeOrder, false, nil); err != nil {
			log.Fatalf("RabbitMq Bind Queue err:%v", err)
		}
		return nil
	}, e.CancelHandler)
	err := consumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RabbitMq Consumer Start!!!")
}

//CancelHandler 订单自动取消消息处理
func (e *Event) CancelHandler(body []byte) error {
	msg := &message.OrderMessage{}
	if err := json.Unmarshal(body, msg); err != nil {
		return err
	}
	if err := e.srv.OrderCancel(context.Background(), msg.MemberID, msg.OrderID); err != nil {
		logger.Warnf("[event] order cancel err: %v", err)
	}
	return nil
}

//Handler 秒杀订单消息处理
func (e *Event) Handler(ctx context.Context, message *pb.Event) error {
	logger.Infof("[event] handler message: %v", message)
	if err := e.srv.SubmitSeckillOrder(ctx, message.MemberId, message.SkuId, message.AddressId,
		int(message.Price), int(message.Num), message.OrderNo); err != nil {
		logger.Warnf("[event] handler message: %v", message)
		return err
	}
	return nil
}
