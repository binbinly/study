package main

import (
	"log"

	"common/app"
	"common/constvar"
	"common/orm"
	pb "common/proto/member"
	"member/cmd"
	"member/conf"
	"member/handler"
	"member/service"
	"pkg/redis"
)

func main() {
	svc := app.NewService("member", conf.Conf,
		app.WithClient(),
		app.WithName(constvar.ServiceMember),
		app.WithMigrate(func() {
			cmd.Migrate()
		}), app.WithAuthFunc(handler.Auth))

	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)

	if err := pb.RegisterMemberHandler(svc.Server(), handler.New(service.New(conf.Conf, svc.Client()))); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
