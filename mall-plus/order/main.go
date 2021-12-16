package main

import (
	"log"

	"go-micro.dev/v4"

	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/order"
	"order/cmd"
	"order/conf"
	"order/event"
	"order/handler"
	"order/service"
	"pkg/redis"
)

func main() {
	svc := app.NewService("order", conf.Conf,
		app.WithName(constvar.ServiceOrder),
		app.WithClient(),
		app.WithMigrate(func() {
			cmd.Migrate()
		}), app.WithAuthFunc(handler.Auth))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// init rabbitmq sub
	s := service.New(conf.Conf, svc.Client())
	event.New(s).Init(conf.Conf.AMQP.Addr)

	if err := pb.RegisterOrderHandler(svc.Server(), handler.New(s)); err != nil {
		log.Fatal(err)
	}
	// register subscriber
	if err := micro.RegisterSubscriber(constvar.TopicOrderSeckill, svc.Server(), event.NewKill(s)); err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
