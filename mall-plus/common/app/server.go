package app

import (
	"log"
	"os"
	"time"

	gclient "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	gserver "github.com/asim/go-micro/plugins/server/grpc/v4"
	"github.com/asim/go-micro/plugins/wrapper/breaker/gobreaker/v4"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v4"
	"github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"common/conf"
	"common/constvar"
	"common/wrap"
	"common/wrap/recovery"
	"common/wrap/tracing"
	"pkg/trace"
)

//NewService 实例化一个微服务对象
func NewService(path string, val interface{}, opts ...Option) micro.Service {
	options := Options{version: "latest"}
	for _, o := range opts {
		o(&options)
	}

	mo := []micro.Option{
		// grpc UnaryInterceptor 不会被调用
		micro.Server(gserver.NewServer()),
		micro.WrapHandler(
			wrap.Auth(options.authFunc),
			wrap.Validator(),
			tracing.NewHandlerWrapper(tracing.WithServiceName(options.name)),
			prometheus.NewHandlerWrapper(),
			ratelimit.NewHandlerWrapper(500),
			recovery.NewHandlerWrapper(),
		),
		micro.Name(options.name),
		micro.Version(options.version),
		micro.RegisterTTL(time.Second * 60),
		micro.RegisterInterval(time.Second * 30),
		micro.Flags(
			&cli.BoolFlag{
				Name:  "migrate",
				Usage: "Execute db migrate",
			},
		),
	}
	if options.client {
		mo = append(mo, []micro.Option{
			micro.Client(gclient.NewClient()),
			micro.WrapClient(
				gobreaker.NewClientWrapper(),
				ratelimit.NewClientWrapper(500),
				tracing.NewClientWrapper(),
			),
		}...)
	}

	// Create service
	srv := micro.NewService(mo...)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			if c.String("registry") == "consul" {
				// 注册consul中心
				micro.Registry(consul.NewRegistry(registry.Addrs(c.String("registry_address"))))
			}
			config := c.String("config")
			if config == "consul" {
				conf.LoadConsul(c.String("registry_address"), path, val)
			} else {
				if config == "" { //默认配置文件
					config = "./default.yaml"
				}
				conf.LoadFile(config, path, val)
			}
			if c.String("tracer_address") != "" {
				// init tracer
				_, err := trace.InitTracerProvider(
					trace.WithServiceName(constvar.ServiceProduct),
					trace.WithEndpoint(c.String("tracer_address")))
				if err != nil {
					log.Fatalf("Failed to init tracer: %v", err)
				}
			}
			//迁移数据库
			if c.Bool("migrate") {
				options.migrate()
				os.Exit(1)
			}
			return nil
		}),
	)

	//优雅关闭服务
	//服务器注销服务并等待处理程序完成执行后再退出
	srv.Server().Init(server.Wait(nil))

	return srv
}
