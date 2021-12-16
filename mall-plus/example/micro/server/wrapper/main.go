package main

import (
	"context"
	"example/micro/server/handler"
	"example/micro/server/subscriber"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/server"
	"log"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %d", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}

func logSubWrapper(fn server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, req server.Message) error {
		log.Printf("[Log Sub Wrapper] Before serving publication topic: %v", req.Topic())
		err := fn(ctx, req)
		log.Printf("[Log Sub Wrapper] After serving publication")
		return err
	}
}

func main() {
	cmd.Init()

	md := server.DefaultOptions().Metadata
	md["datacenter"] = "local"

	server.DefaultServer = server.NewServer(
		server.WrapHandler(logWrapper),
		server.WrapSubscriber(logSubWrapper),
		server.Metadata(md))

	server.Init(
		server.Name("go.micro.srv.example"))

	server.Handle(
		server.NewHandler(
			new(handler.Example)))

	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.example",
			new(subscriber.Example))); err != nil {
		log.Fatal(err)
	}

	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.example",
			subscriber.Handler)); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
