package main

import (
	"example/micro/cache/handler"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	pb "example/micro/cache/proto"
)

var (
	service = "go.micro.srv.cache"
	version = "latest"
)

func main() {
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version))
	srv.Init()

	pb.RegisterCacheHandler(srv.Server(), handler.NewCache())

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
