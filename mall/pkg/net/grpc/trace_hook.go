package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"mall/pkg/log"
)

var _ Hook = (*TracingHook)(nil)

//TracingHook 链路追踪钩子
type TracingHook struct {
	tracer opentracing.Tracer
}

//NewTracingHook 实例化
func NewTracingHook() *TracingHook {
	return &TracingHook{tracer: opentracing.GlobalTracer()}
}

// BeforeHandler 执行操作前调用
func (t TracingHook) BeforeHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	spanContext, err := t.tracer.Extract(
		opentracing.TextMap,
		MDCarrier{MD: md},
	)
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		log.Warnf("[grpc.hook] extract from metadata err: %v", err)
	} else {
		span := t.tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(spanContext),
			ext.SpanKindRPCServer,
		)
		ext.Component.Set(span, "gRPC Server")

		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return ctx, nil
}

//AfterHandler 执行操作后调用
func (t TracingHook) AfterHandler(ctx context.Context, info *grpc.UnaryServerInfo, err error) error {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return err
	}
	if err != nil {
		ext.LogError(span, err)
	}
	span.Finish()
	return err
}
