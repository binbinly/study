package main

import (
	"context"
	example "example/micro/server/proto/example"
	"fmt"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type firstNodeSelector struct {
	opts selector.Options
}

func (m *firstNodeSelector) Init(opts ...selector.Option) error {
	for _, opt := range opts {
		opt(&m.opts)
	}
	return nil
}

func (m *firstNodeSelector) Options() selector.Options {
	return m.opts
}

func (m *firstNodeSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	services, err := m.opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var sopts selector.SelectOptions
	for _, opt := range opts {
		opt(&sopts)
	}

	for _, filter := range sopts.Filters {
		services = filter(services)
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	if len(services[0].Nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	return func() (*registry.Node, error) {
		return services[0].Nodes[0], nil
	}, nil
}

func (n *firstNodeSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *firstNodeSelector) Reset(service string) {
	return
}

func (n *firstNodeSelector) Close() error {
	return nil
}

func (n *firstNodeSelector) String() string {
	return "first"
}

func FirstNodeSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &firstNodeSelector{sopts}
}

func call(i int) {
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	rsp := &example.Response{}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err:", err, rsp)
		return
	}

	fmt.Println("call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Selector(FirstNodeSelector()))

	fmt.Println("\n --- Call Example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
