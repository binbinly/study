package main

import (
	"context"
	example "example/micro/server/proto/example"
	"fmt"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
	"time"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

type traceWrapper struct {
	client.Client
}

func (t *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctx = metadata.NewContext(ctx, map[string]string{
		"X-Trace-Id": fmt.Sprintf("%d", time.Now().Unix()),
	})
	return t.Client.Call(ctx, req, rsp)
}

func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func traceWrap(c client.Client) client.Client {
	return &traceWrapper{c}
}

func metricsWrap(cf client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := cf(ctx, node, req, rsp, opts)
		fmt.Printf("[Metrics Wwrapper] called: %v %s.%s duration: %v\n", node, req.Service(), req.Endpoint(), time.Since(t))
		return err
	}
}

func call(i int) {
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "john",
	})

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User_id": "john",
		"X-From-Id": "script",
	})

	rsp := &example.Response{}

	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err:", err, rsp)
		return
	}

	fmt.Println("call:", i, "rsp", rsp.Msg)
}

func main() {
	cmd.Init()

	fmt.Println("\n --- Log Wrapper example ---")

	client.DefaultClient = logWrap(client.DefaultClient)

	call(0)

	fmt.Println("\n --- Log+Trace Wrapper example ---")

	client.DefaultClient = client.NewClient(
		client.Wrap(traceWrap),
		client.Wrap(logWrap))

	call(1)

	fmt.Println("\n --- Metrics Wrapper example ---")

	client.DefaultClient = client.NewClient(
		client.WrapCall(metricsWrap))

	call(2)
}
