package main

import (
	"log"

	"common/app"
	"common/constvar"
	pb "common/proto/seckill"
	"pkg/redis"
	"seckill/conf"
	"seckill/handler"
	"seckill/service"
)

func main() {
	svc := app.NewService("seckill", conf.Conf,
		app.WithName(constvar.ServiceSeckill),
		app.WithAuthFunc(handler.Auth))

	// init redis
	redis.Init(&conf.Conf.Redis)

	s := service.New(conf.Conf, svc.Client())
	if err := pb.RegisterSeckillHandler(svc.Server(), handler.New(s)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
