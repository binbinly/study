package logic

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"chat-micro/app/constvar"
	"chat-micro/app/logic"
	"chat-micro/app/logic/conf"
	handler "chat-micro/app/logic/handler/grpc"
	"chat-micro/app/logic/server"
	"chat-micro/app/logic/service"
	"chat-micro/internal/orm"
	"chat-micro/pkg/broker"
	"chat-micro/pkg/broker/nats"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/minio"
	"chat-micro/pkg/redis"
	"chat-micro/pkg/registry"
	"chat-micro/pkg/registry/consul"
	"chat-micro/pkg/trace"
	pb "chat-micro/proto/logic"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "logic",
		Short:        "Start chat-micro logic server",
		Example:      "chat-micro logic -c config/logic/default.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/logic/default.yml", "Start server with provided configuration file")
}

func setup() {
	// init conf
	conf.Init(configYml)
	// init log
	l, err := logger.NewZapLogger(logger.WithLevel(logger.WarnLevel))
	if err != nil {
		log.Fatalf("Failed to init zap log err: %v", err)
	}
	if conf.Conf.Debug {
		l.Init(logger.WithLevel(logger.DebugLevel))
	}
	logger.DefaultLogger = l

	// init tracer
	_, err = trace.InitTracerProvider(
		trace.WithServiceName(constvar.ServiceLogic),
		trace.WithEndpoint(conf.Conf.TracerAddress))
	if err != nil {
		log.Fatalf("Failed to init tracer: %v", err)
	}
}

//run 核心业务服务启动
func run() {
	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// init minio storage
	storage := minio.New(&conf.Conf.Minio)
	if err := storage.Init(); err != nil {
		log.Fatalf("Failed to init minio storage: %v", err)
	}
	b := nats.NewBroker(broker.Addrs(conf.Conf.BrokerAddress))
	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	grpcServer := server.NewGRPCServer(&conf.Conf.GRPC)
	app := logic.New(
		logic.WithName(constvar.ServiceLogic),
		logic.WithRegistry(
			consul.NewRegistry(
				registry.Addrs(conf.Conf.RegistryAddress),
				consul.TCPCheck(time.Second*30),
			),
		),
		logic.WithServer(grpcServer, server.NewHTTPServer(&conf.Conf.HTTP)),
	)

	//注册处理器
	pb.RegisterLogicServer(grpcServer, handler.New(service.New(
		service.WithName(constvar.ServiceLogic),
		service.WithTopic(constvar.TaskTopic),
		service.WithBroker(b),
		service.WithJwtSecret(conf.Conf.JwtSecret),
		service.WithStorage(storage))))

	//start server
	if err := app.Run(); err != nil {
		log.Fatalf("App Run err: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Server is exiting") // close http server
			if err := app.Stop(); err != nil {
				log.Printf("Failed to stop app err: %v", err)
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
