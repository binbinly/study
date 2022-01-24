package grpc

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"chat-micro/pkg/util"
)

// Server is a gRPC server wrapper.
// nolint
type Server struct {
	*grpc.Server
	opts     options
	ctx      context.Context
	lis      net.Listener
	once     sync.Once
	endpoint *url.URL
	health   *health.Server
}

// NewServer creates a gRPC server by options.
func NewServer(opts ...Option) *Server {
	srv := &Server{
		opts:   newOptions(opts...),
		health: health.NewServer(),
	}

	gopts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(srv.opts.maxMsgSize),
		grpc.MaxSendMsgSize(srv.opts.maxMsgSize),
	}
	if srv.opts.keepalive {
		gopts = append(gopts, grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Second, //如果客户端闲置 15 秒，发送GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,  //在强行关闭连接之前，等待 5 秒钟，以使挂起的RPC完成
			MaxConnectionAge:      30 * time.Second, //处于活动状态的连接超过 30 秒，发送GOAWAY
			Time:                  5 * time.Second,  //心跳时间 5 秒
			Timeout:               time.Second,      //连接超时前等待 1 秒的 ping ack
		}))
	}
	if len(srv.opts.grpcOpts) > 0 {
		gopts = append(gopts, srv.opts.grpcOpts...)
	}

	grpcServer := grpc.NewServer(gopts...)

	// see https://github.com/grpc/grpc/blob/master/doc/health-checking.md for more
	srv.health.SetServingStatus("", healthPb.HealthCheckResponse_SERVING)
	healthPb.RegisterHealthServer(grpcServer, srv.health)
	reflection.Register(grpcServer)

	srv.Server = grpcServer

	return srv
}

// Endpoint return a real address to registry endpoint.
// examples:
//   grpc://127.0.0.1:9090
func (s *Server) Endpoint() (*url.URL, error) {
	addr, err := util.Extract(s.opts.address, s.lis)
	if err != nil {
		return nil, err
	}
	s.endpoint = &url.URL{Scheme: "grpc", Host: addr}
	return s.endpoint, nil
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.opts.network, s.opts.address)
	if err != nil {
		return err
	}
	s.lis = lis

	if _, err := s.Endpoint(); err != nil {
		return err
	}

	s.ctx = ctx
	log.Printf("[gRPC] server is listening on: %s", s.lis.Addr().String())
	return s.Serve(s.lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	log.Printf("[gRPC] server is stopping")
	return nil
}
