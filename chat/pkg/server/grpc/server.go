package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

var srv *Server

// ServerConfig is GRPC server config.
type ServerConfig struct {
	Network           string
	Port              int
	Timeout           time.Duration
	IdleTimeout       time.Duration //如果客户端闲置 x 秒，发送GOAWAY
	MaxLifeTime       time.Duration //如果任何连接仍然存在超过 x 秒，发送GOAWAY
	ForceCloseWait    time.Duration //在强行关闭连接之前，等待 x 秒钟，以使挂起的RPC完成
	KeepAliveInterval time.Duration //如果客户端闲置 x 秒钟，对其进行ping操作，以确保连接仍处于活动状态
	KeepAliveTimeout  time.Duration //假设连接中断，等待 x 秒钟以进行ping确认
}

type Server struct {
	Srv *grpc.Server
	//该Server的请求开始时Hook函数
	OnRequestStart func(ctx context.Context, req interface{})
	//该Server的请求结束时的Hook函数
	OnRequestEnd func(ctx context.Context, req interface{})
}

func NewServer(c *ServerConfig, host string, reg func(srv *grpc.Server)) *Server {
	srv = &Server{
		Srv: New(c, host, reg),
	}
	return srv
}

//SetOnRequestStart 该Server的请求开始时Hook函数
func (s *Server) SetOnRequestStart(hook func(ctx context.Context, req interface{})) {
	s.OnRequestStart = hook
}

//SetOnRequestEnd 该Server的请求结束时Hook函数
func (s *Server) SetOnRequestEnd(hook func(ctx context.Context, req interface{})) {
	s.OnRequestEnd = hook
}

//CallOnRequestStart 调用连接OnRequestStart Hook函数
func (s *Server) CallOnRequestStart(ctx context.Context, req interface{}) {
	if s.OnRequestStart != nil {
		s.OnRequestStart(ctx, req)
	}
}

//CallOnRequestEnd 调用连接OnRequestEnd Hook函数
func (s *Server) CallOnRequestEnd(ctx context.Context, req interface{}) {
	if s.OnRequestEnd != nil {
		s.OnRequestEnd(ctx, req)
	}
}

//Stop 停止服务
func (s *Server) Stop()  {
	s.Srv.GracefulStop()
}

// New grpc server
func New(c *ServerConfig, host string, reg func(srv *grpc.Server)) *grpc.Server {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     c.IdleTimeout,
		MaxConnectionAgeGrace: c.ForceCloseWait,
		Time:                  c.KeepAliveInterval,
		Timeout:               c.KeepAliveTimeout,
		MaxConnectionAge:      c.MaxLifeTime,
	})
	s := grpc.NewServer(keepParams, grpc.UnaryInterceptor(UnaryServerInterceptor))
	// 注册rpc服务类
	reg(s)
	healthServer := &HealthImpl{}
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	addr := fmt.Sprintf("%s:%d", host, c.Port)
		lis, err := net.Listen(c.Network, addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if e := s.Serve(lis); e != nil {
			log.Panicf("failed to serve grpc server: %v", e)
		}
	}()
	log.Printf("serve grpc server is success, addr:%s", addr)
	return s
}

// UnaryServerInterceptor server拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch info.Server.(type) {
	case *HealthImpl: //consul健康检查
		return handler(ctx, req)
	default:
		srv.OnRequestStart(ctx, req)
		resp, err := handler(ctx, req)
		srv.OnRequestEnd(ctx, req)
		return resp, err
	}
}
