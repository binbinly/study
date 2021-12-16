package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"

	hello "example/micro/gateway/proto/hello"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel :=context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main()  {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}