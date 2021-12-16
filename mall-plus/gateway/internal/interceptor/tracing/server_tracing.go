package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Unary RPCs.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := options{
		TracerProvider: otel.GetTracerProvider(),
		Propagators:    otel.GetTextMapPropagator(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	tracer := o.TracerProvider.Tracer(
		o.ServiceName,
	)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = make(metadata.MD)
		}
		ctx = o.Propagators.Extract(ctx, MetadataCarrier(md))

		ctx, span := tracer.Start(ctx,
			info.FullMethod,
		)
		defer span.End()

		resp, err := handler(ctx, req)
		return resp, err
	}
}