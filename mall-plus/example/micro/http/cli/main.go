package main

import (
	"context"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"log"
)

func main()  {
	CallHttpServer()
}

func CallHttpServer()  {
	r := registry.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := httpClient.NewClient(client.Selector(s))
	request := c.NewRequest("demo-http", "/demo", "", client.WithContentType("application/json"))

	response := new(map[string]interface{})

	err := c.Call(context.Background(), request, response)
	log.Printf("err: %v response: %#v\n", err, response)
}
