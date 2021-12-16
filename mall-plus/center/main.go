package main

import (
	"log"

	"center/cmd"
	"center/conf"
	"center/handler"
	"center/service"
	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/center"
	"pkg/redis"
)

func main() {
	svc := app.NewService("center", conf.Conf,
		app.WithName(constvar.ServiceCenter),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterUserHandler(svc.Server(), handler.New(service.New(conf.Conf))); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
