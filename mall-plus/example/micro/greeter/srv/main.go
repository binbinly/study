package main

import (
	"context"
	hello "example/micro/greeter/srv/proto/hello"
	"example/grpc"
	"go-micro.dev/v4"
	"log"
	"time"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Println("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	time.Sleep(time.Second * 3)
	log.Println("Received Say.Hello request 1")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Server(grpc.NewServer()),
		)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
