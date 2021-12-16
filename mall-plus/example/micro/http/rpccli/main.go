package main

import (
	"context"
	hello "example/micro/greeter/srv/proto/hello"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Registry(registry.NewRegistry()))

	service.Init()
	c := service.Client()

	req := &hello.Request{Name: "call grpc server by http client"}
	request := client.NewRequest("go.micro.srv.greeter", "Say.Hello", req)

	response := new(hello.Response)
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)
}
