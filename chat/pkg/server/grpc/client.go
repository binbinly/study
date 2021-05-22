package grpc

import (
	"context"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"chat/pkg/log"
	"chat/pkg/server/grpc/meta"
)

// ClientConfig is GRPC client config.
type ClientConfig struct {
	Timeout          time.Duration //请求超时
	KeepAliveTime    time.Duration //如果客户端闲置 x 秒钟，对其进行ping操作，以确保连接仍处于活动状态
	KeepAliveTimeout time.Duration //假设连接中断，等待 x 秒钟以进行ping确认
}

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
)

// NewRpcClientConn 创建一个 grpc客户端连接
func NewRpcClientConn(c *ClientConfig, ctx context.Context, target string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, target,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(UnaryClientInterceptor),
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                c.KeepAliveTime,
				Timeout:             c.KeepAliveTimeout,
				PermitWithoutStream: true, // 即使没有活跃的流也发送ping
			}),
		}...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// UnaryClientInterceptor 拦截器
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Infof("before invoker. method: %v, request:%v", method, req)
	rpcMeta := meta.GetRpcMeta(ctx)
	// 熔断器
	err := hystrix.Do(rpcMeta.ServiceName, func() (err error) {
		err = invoker(ctx, method, req, reply, cc, opts...)
		return err
	}, func(err error) error {
		log.Warnf("hystrix err:%v", err)
		return err
	})
	log.Infof("after invoker. reply: %v", reply)
	return err
}
