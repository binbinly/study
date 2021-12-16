package app

import (
	"log"

	"gateway/internal/interceptor"
	"gateway/internal/interceptor/metrics"
	"gateway/internal/interceptor/sentinel"
	"gateway/internal/interceptor/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
)

//newRPCClientConn 实例化一个 grpc客户端连接
func (a *App) newRPCClientConn(service *ServiceItem) *grpc.ClientConn {
	target := "consul:///" + service.Name
	conn, err := grpc.DialContext(a.ctx, target,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(
				interceptor.ValidatorClientInterceptor(),
				interceptor.TimeoutClientInterceptor(service.Timeout),
				metrics.UnaryClientInterceptor(metrics.NewClientMetrics()),
				tracing.UnaryClientInterceptor(tracing.WithServiceName(service.Name)),
				sentinel.FlowClientInterceptor(
					sentinel.WithResource(service.Name),
					sentinel.WithThreshold(service.QPSLimit)),
				sentinel.CircuitBreakerClientInterceptor(
					sentinel.WithBreakerResource(service.Name)),
			),
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				PermitWithoutStream: true, // 即使没有活跃的流也发送ping
			}),
		}...)
	if err != nil {
		log.Fatalf("failed new grpc client err: %v by target: %v", err, target)
	}
	return conn
}
