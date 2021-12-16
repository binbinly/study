package main

import (
	"example/micro/server/subscriber"
	"log"

	"context"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/server"

	example "example/micro/server/proto/example"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Println("Received Example.Call request")
	rsp.Msg = server.DefaultOptions().Id + ": Hello" + req.Name
	return nil
}

func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Printf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Printf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func main() {
	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.example"))

	server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.example",
			new(subscriber.Example)))

	server.Subscribe(
		server.NewSubscriber(
			"topic.go.micro.srv.example",
			subscriber.Handler))

	example.RegisterExampleHandler(server.DefaultServer, new(Example))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
