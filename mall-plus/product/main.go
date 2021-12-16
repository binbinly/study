package main

import (
	"log"

	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/product"
	"pkg/redis"
	"product/cmd"
	"product/conf"
	"product/handler"
	"product/service"
)

func main() {
	svc := app.NewService("product", conf.Conf,
		app.WithName(constvar.ServiceProduct),
		app.WithClient(),
		app.WithMigrate(func() {
			cmd.Migrate()
		}), app.WithAuthFunc(handler.Auth))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterProductHandler(svc.Server(), handler.New(service.New(conf.Conf, svc.Client()))); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
