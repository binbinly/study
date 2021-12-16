package main

import (
	"log"

	"common/app"
	"common/constvar"
	pb "common/proto/third"
	"pkg/redis"
	"third-party/conf"
	"third-party/handler"
	"third-party/service"
)

func main() {
	svc := app.NewService("third", conf.Conf,
		app.WithName(constvar.ServiceThird))

	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterThirdHandler(svc.Server(), handler.New(service.New(conf.Conf))); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
