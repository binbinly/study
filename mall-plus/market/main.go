package main

import (
	"log"

	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/market"
	"market/cmd"
	"market/conf"
	"market/handler"
	"market/service"
	"pkg/redis"
)

func main() {
	svc := app.NewService("market", conf.Conf,
		app.WithName(constvar.ServiceMarket),
		app.WithMigrate(func() {
			cmd.Migrate()
		}), app.WithAuthFunc(handler.Auth))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterMarketHandler(svc.Server(), handler.New(service.New(conf.Conf))); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
