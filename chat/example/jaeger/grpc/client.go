package main

import (
	"chat/example/grpc/hello"
	"chat/example/jaeger/grpc/cerrier"
	"chat/example/jaeger/lib"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	tracer, closer := lib.Init("Hello-World")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", "gopher")
	defer span.Finish()

	serviceAddress := "127.0.0.1:50051"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, serviceAddress, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithUnaryInterceptor(ClientInterceptor(tracer)))
	if err != nil {
		fmt.Printf("grpc dial err: %v", err)
		panic("grpc dial err")
	}
	defer func() {
		_ = conn.Close()
	}()

	userClient := hello.NewHelloServiceClient(conn)
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req := &hello.HelloRequest{
		Name: "gopher",
	}

	reply, err := userClient.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Printf("reply : %+v", reply)
}

// ClientInterceptor 客户端拦截器
// https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor
func ClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, request, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		//一个RPC调用的服务端的span，和RPC服务客户端的span构成ChildOf关系
		//var parentCtx opentracing.SpanContext
		//parentSpan := opentracing.SpanFromContext(ctx)
		//if parentSpan != nil {
		//	parentCtx = parentSpan.Context()
		//}
		//span := tracer.StartSpan(
		//	method,
		//	opentracing.ChildOf(parentCtx),
		//)
		span, ctx := opentracing.StartSpanFromContext(ctx, method)
		defer span.Finish()

		ext.Component.Set(span, "gRPC Client")
		ext.SpanKindRPCClient.Set(span)

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			cerrier.MDCarrier{md}, // 自定义 carrier
		)

		if err != nil {
			log.Fatalf("inject span error :%v", err.Error())
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, request, reply, cc, opts...)

		if err != nil {
			log.Fatalf("call error : %v", err.Error())
		}
		return err
	}
}
