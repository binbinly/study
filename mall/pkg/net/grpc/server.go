package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"mall/pkg/app"
	"mall/pkg/registry"
)

//Hook 请求前后执行钩子
type Hook interface {
	// 执行前调用
	BeforeHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (context.Context, error)
	// 执行后调用
	AfterHandler(ctx context.Context, info *grpc.UnaryServerInfo, err error) error
}

// ServerConfig is GRPC server config.
type ServerConfig struct {
	Network           string
	Port              int
	QPSLimit          int //服务限流 默认0，不限制
	Timeout           time.Duration
	IdleTimeout       time.Duration //如果客户端闲置 x 秒，发送GOAWAY
	MaxLifeTime       time.Duration //如果任何连接仍然存在超过 x 秒，发送GOAWAY
	ForceCloseWait    time.Duration //在强行关闭连接之前，等待 x 秒钟，以使挂起的RPC完成
	KeepAliveInterval time.Duration //如果客户端闲置 x 秒钟，对其进行ping操作，以确保连接仍处于活动状态
	KeepAliveTimeout  time.Duration //假设连接中断，等待 x 秒钟以进行ping确认
}

//Server grpc服务结构
type Server struct {
	*grpc.Server //服务器

	c       *ServerConfig //服务配置
	limiter *rate.Limiter //令牌桶限流
	hooks   []Hook        //该Server的请求前后Hook函数
	options *app.Options
}

//NewServer 实例化服务器
func NewServer(c *ServerConfig, opts ...app.Option) *Server {
	s := &Server{
		c: c,
		options: &app.Options{},
	}
	for _, o := range opts {
		o(s.options)
	}
	if s.c.QPSLimit > 0 {
		s.limiter = rate.NewLimiter(rate.Limit(s.c.QPSLimit), s.c.QPSLimit)
	}
	return s
}

//Init 初始化GRPC服务
func (s *Server) Init() {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     s.c.IdleTimeout,
		MaxConnectionAgeGrace: s.c.ForceCloseWait,
		Time:                  s.c.KeepAliveInterval,
		Timeout:               s.c.KeepAliveTimeout,
		MaxConnectionAge:      s.c.MaxLifeTime,
	})
	s.Server = grpc.NewServer(keepParams, grpc.UnaryInterceptor(s.interceptor))

	//健康检查
	healthServer := &HealthImpl{}
	grpc_health_v1.RegisterHealthServer(s.Server, healthServer)
}

//Start 启动服务
func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%d", s.options.Host, s.c.Port)
	lis, err := net.Listen(s.c.Network, addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err = s.Server.Serve(lis); err != nil {
			log.Panicf("failed to serve grpc server: %v", err)
		}
	}()
	log.Printf("server grpc start is success, addr:%s", addr)
	s.register()
}

//GetServerID 获取服务器ID
func (s *Server) GetServerID() string {
	return app.BuildServerID(s.options.ID, s.c.Port)
}

//register 注册服务
func (s *Server) register() {
	if s.options.Register == nil {
		return
	}
	err := s.options.Register.Register(context.Background(), &registry.Service{
		ID:   app.BuildServerID(s.options.ID, s.c.Port),
		Name: s.options.Name,
		IP:   s.options.Host,
		Port: s.c.Port,
		Check: registry.Check{
			GRPC: fmt.Sprintf("%v:%v/%v", s.options.Host, s.c.Port, s.options.Name),
		},
	})
	if err != nil {
		log.Fatalf("failed to serve grpc register %s server: %v", s.options.Name, err)
	}
	log.Printf("server grpc register success id:%v, name:%v", app.BuildServerID(s.options.ID, s.c.Port), s.options.Name)
}

//AddHook 添加钩子
func (s *Server) AddHook(hook ...Hook) {
	s.hooks = append(s.hooks, hook...)
}

//Stop 停止服务
func (s *Server) Stop() {
	s.Server.GracefulStop()
	if s.options.Register != nil { // 服务注销
		s.options.Register.Unregister(context.Background(), &registry.Service{ID: app.BuildServerID(s.options.ID, s.c.Port)})
	}
}

//interceptor 拦截器
func (s *Server) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
	switch info.Server.(type) {
	case *HealthImpl: //consul健康检查
		return handler(ctx, req)
	default:
		// 触发限流
		if !s.limiter.Allow() {
			err = status.Error(codes.ResourceExhausted, "rate limited")
			return
		}
		var hookIndex int

		for ; hookIndex < len(s.hooks); hookIndex++ {
			ctx, err = s.hooks[hookIndex].BeforeHandler(ctx, req, info)
		}
		// 执行操作
		if err == nil {
			rsp, err = handler(ctx, req)
		}
		for hookIndex--; hookIndex >= 0; hookIndex-- {
			err = s.hooks[hookIndex].AfterHandler(ctx, info, err)
		}
		return rsp, err
	}
}
