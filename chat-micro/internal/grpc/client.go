package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"chat-micro/pkg/logger"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
)

//NewClientConn 实例化一个 grpc客户端连接
func NewClientConn(ctx context.Context, target string) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, target,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(),
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				PermitWithoutStream: true, // 即使没有活跃的流也发送ping
			}),
		}...)
	if err != nil {
		logger.Fatalf("failed new grpc client err: %v by target: %v", err, target)
	}
	return conn
}
