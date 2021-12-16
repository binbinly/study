package main

import (
	"log"

	"cart/conf"
	"cart/handler"
	"cart/service"
	"common/app"
	"common/constvar"
	pb "common/proto/cart"
	"pkg/redis"
)

func main() {
	svc := app.NewService("cart", conf.Conf,
		app.WithName(constvar.ServiceCart),
		app.WithClient(),
		app.WithAuthFunc(handler.Auth))

	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterCartHandler(svc.Server(), handler.New(service.New(conf.Conf, svc.Client()))); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
