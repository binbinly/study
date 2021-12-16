package tracing

import (
	"context"
	"go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"pkg/net/ip"
	"pkg/utils"
)

//see: https://goframe.org/pages/viewpage.action?pageId=3673684
const (
	tracingName                     = "grpc-client"
	tracingMaxContentLogSize        = 256 * 1024 // Max log size for request and response body.
	tracingEventGrpcRequest         = "grpc.request"
	tracingEventGrpcRequestMessage  = "grpc.request.message"
	tracingEventGrpcResponse        = "grpc.response"
	tracingEventGrpcResponseMessage = "grpc.response.message"
)

var (
	localIP = ip.GetLocalIP()
)

type options struct {
	ServiceName    string
	TracerProvider trace.TracerProvider
	Propagators    propagation.TextMapPropagator
}

// Option specifies instrumentation configuration options.
type Option func(*options)

//WithServiceName set service name
func WithServiceName(name string) Option {
	return func(o *options) {
		o.ServiceName = name
	}
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.Propagators = propagators
	}
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(o *options) {
		o.TracerProvider = provider
	}
}

// UnaryClientInterceptor returns a new unary client interceptor for OpenTracing.
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	o := options{
		ServiceName:    tracingName,
		TracerProvider: otel.GetTracerProvider(),
		Propagators:    otel.GetTextMapPropagator(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	tracer := o.TracerProvider.Tracer(
		o.ServiceName,
		trace.WithInstrumentationVersion(contrib.SemVersion()),
	)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var span trace.Span
		ctx, span = tracer.Start(ctx, method, []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(
				semconv.RPCServiceKey.String(cc.Target()),
				semconv.RPCSystemKey.String("grpc"),
			),
		}...)
		defer span.End()

		md, _ := metadata.FromOutgoingContext(ctx)
		mdCopy := md.Copy()

		o.Propagators.Inject(ctx, MetadataCarrier(mdCopy))

		span.SetAttributes(
			attribute.String("hostname", utils.Hostname),
			attribute.String("local-ip", localIP))

		span.AddEvent(tracingEventGrpcRequest, trace.WithAttributes(
			attribute.String(
				tracingEventGrpcRequestMessage,
				MarshalMessageToJSONStringForTracing(
					req, "Request", tracingMaxContentLogSize,
				),
			),
		))

		//链路关键
		ctx = metadata.NewOutgoingContext(ctx, mdCopy)
		err := invoker(ctx, method, req, reply, cc, opts...)

		span.AddEvent(tracingEventGrpcResponse, trace.WithAttributes(
			attribute.String(
				tracingEventGrpcResponseMessage,
				MarshalMessageToJSONStringForTracing(
					reply, "Response", tracingMaxContentLogSize,
				),
			),
		))
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(semconv.RPCGRPCStatusCodeKey.String(s.Code().String()))
		} else {
			span.SetStatus(codes.Ok, "OK")
		}
		return err
	}
}
