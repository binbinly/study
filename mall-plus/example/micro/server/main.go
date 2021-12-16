package main

import (
	"example/micro/server/handler"
	"example/micro/server/subscriber"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/server"
	"log"
)

func main() {
	cmd.Init()

	server.Init(server.Name("go.micro.srv.example"))

	server.Handle(server.NewHandler(new(handler.Example)))

	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.example",
			new(subscriber.Example))); err != nil {
		log.Fatal(err)
	}

	if err := server.Subscribe(
		server.NewSubscriber("topic.example", subscriber.Handler)); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
