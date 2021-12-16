package main

import (
	"log"

	"go-micro.dev/v4"

	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/warehouse"
	"pkg/redis"
	"warehouse/cmd"
	"warehouse/conf"
	"warehouse/event"
	"warehouse/handler"
	"warehouse/service"
)

func main() {
	svc := app.NewService("warehouse", conf.Conf,
		app.WithName(constvar.ServiceWarehouse),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)

	s := service.New(conf.Conf)
	if err := pb.RegisterWarehouseHandler(svc.Server(), handler.New(s)); err != nil {
		log.Fatal(err)
	}
	// register subscriber
	if err := micro.RegisterSubscriber(constvar.TopicWarehouse, svc.Server(), event.New(s)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
