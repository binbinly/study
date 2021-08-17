package grpc

import (
	"context"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	"chat/pkg/log"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
)

// ClientConfig is GRPC client config.
type ClientConfig struct {
	ServiceName      string        //服务名
	QPSLimit         int           //并发限制
	Timeout          time.Duration //请求超时
	KeepAliveTime    time.Duration //如果客户端闲置 x 秒钟，对其进行ping操作，以确保连接仍处于活动状态
	KeepAliveTimeout time.Duration //假设连接中断，等待 x 秒钟以进行ping确认
}

// NewRPCClientConn 实例化一个 grpc客户端连接
func NewRPCClientConn(c *ClientConfig, target string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(ClientInterceptor(c.ServiceName)),
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
		log.Fatalf("failed new grpc client conn: %v", err)
	}
	// 初始化熔断器配置
	hystrix.ConfigureCommand(c.ServiceName, hystrix.CommandConfig{
		Timeout:                1000,       // 超时配置，默认1000ms
		MaxConcurrentRequests:  c.QPSLimit, // 并发控制，默认是10
		SleepWindow:            5000,       // 熔断器打开之后，冷却的时间，默认是5000ms
		RequestVolumeThreshold: 20,         // 一个统计窗口的请求数量，默认是20
		ErrorPercentThreshold:  50,         // 失败百分比，默认是50%
	})
	return conn
}

// ClientInterceptor 拦截器
func ClientInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Infof("before invoker. method: %v, request:%v", method, req)
		// 链路追踪
		tracer := opentracing.GlobalTracer()
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
			MDCarrier{MD: md}, // 自定义 carrier
		)

		if err != nil {
			log.Warnf("[grpc.client] trace inject span error :%v", err)
		} else {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		// 熔断器
		// TODO 第二个参数熔断时触发，暂时不做任何处理
		err = hystrix.Do(serviceName, func() (err error) {
			err = invoker(ctx, method, req, reply, cc, opts...)
			return err
		}, nil)
		if err != nil {
			ext.LogError(span, err)
		}
		return err
	}
}
