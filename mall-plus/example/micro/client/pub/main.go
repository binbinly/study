package main

import (
	"context"
	example "example/micro/server/proto/example"
	"fmt"
	"go-micro.dev/v4"
)

func pub(i int, p micro.Publisher) {
	msg := &example.Message{
		Say: fmt.Sprintf("This is on async message: %d", i),
	}

	if err := p.Publish(context.TODO(), msg); err != nil {
		fmt.Println("pub err:", err)
		return
	}

	fmt.Printf("Published %d: %v\n", i, msg)
}

func main() {
	service := micro.NewService()
	service.Init()

	p := micro.NewPublisher("example", service.Client())

	fmt.Println("\n --- Publisher example ---")

	for i := 0; i < 10; i++ {
		pub(i, p)
	}
}
