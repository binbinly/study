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
	"sync"
	"time"
)

type dcSelector struct {
	opts selector.Options
}

var (
	datacenter = "local"
)

func init()  {
	rand.Seed(time.Now().Unix())
}

func (n *dcSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *dcSelector) Options() selector.Options {
	return n.opts
}

func (n *dcSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	services, err := n.opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var nodes []*registry.Node

	for _, s := range services {
		for _, node := range s.Nodes {
			if node.Metadata["datacenter"] == datacenter {
				nodes = append(nodes, node)
			}
		}
	}

	if len(nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	var i int
	var mtx sync.Mutex

	return func() (*registry.Node, error) {
		mtx.Lock()
		defer mtx.Unlock()
		i++
		return nodes[i%len(nodes)], nil
	}, nil
}
func (n *dcSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *dcSelector) Reset(service string) {
	return
}

func (n *dcSelector) Close() error {
	return nil
}

func (n *dcSelector) String() string {
	return "dc"
}

func DCSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &dcSelector{sopts}
}

func call(i int)  {
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	rsp := &example.Response{}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err:", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}


func main()  {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Selector(DCSelector()))

	fmt.Println("\n --- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
