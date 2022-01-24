package connect

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"chat-micro/app/connect"
	"chat-micro/app/connect/conf"
	handler "chat-micro/app/connect/handler/grpc"
	"chat-micro/app/connect/server"
	"chat-micro/app/constvar"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/registry"
	"chat-micro/pkg/registry/consul"
	"chat-micro/pkg/trace"
	pb "chat-micro/proto/connect"
)

const (
	registryAddress = "registry_address"
	tracerAddress   = "tracer_address"
	wsAddress       = "ws_address"
	tcpAddress      = "tcp_address"
	rpcAddress      = "rpc_address"
)

var (
	configYml string
	ws        bool
	tcp       bool
	StartCmd  = &cobra.Command{
		Use:          "connect",
		Short:        "Start chat-micro connect server",
		Example:      "chat-micro connect -c config/connect/default.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			//flag()
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/connect/default.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVar(&ws, "ws", true, "Start Websocket Server")
	StartCmd.PersistentFlags().BoolVar(&tcp, "tcp", false, "Start TCP Server")
	StartCmd.PersistentFlags().Bool("dev", false, "Enable Debug [$CHAT_DEBUG]")
	StartCmd.PersistentFlags().StringP(wsAddress, "w", "", "Set Websocket Server Address [$CHAT_WS_ADDRESS]")
	StartCmd.PersistentFlags().StringP(tcpAddress, "t", "", "Set TCP Server Address [$CHAT_TCP_ADDRESS]")
	StartCmd.PersistentFlags().StringP(rpcAddress, "r", "", "Set GRPC Server Address [$CHAT_RPC_ADDRESS]")
	StartCmd.PersistentFlags().String(registryAddress, "", "Set registry Address [$CHAT_REGISTRY_ADDRESS]")
	StartCmd.PersistentFlags().String(tracerAddress, "", "Set tracer Address [$CHAT_TRACER_ADDRESS]")

	viper.BindPFlag("registryAddress", StartCmd.PersistentFlags().Lookup(registryAddress))
	viper.BindPFlag("tracerAddress", StartCmd.PersistentFlags().Lookup(tracerAddress))
	viper.BindPFlag("wsAddress", StartCmd.PersistentFlags().Lookup(wsAddress))
	viper.BindPFlag("tcpAddress", StartCmd.PersistentFlags().Lookup(tcpAddress))
	viper.BindPFlag("grpc.Addr", StartCmd.PersistentFlags().Lookup(rpcAddress))
	viper.BindPFlag("debug", StartCmd.PersistentFlags().Lookup("dev"))

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CHAT") // 读取环境变量的前缀为 chat
	viper.BindEnv("wsAddress", "CHAT_WS_ADDRESS")
	viper.BindEnv("tcpAddress", "CHAT_TCP_ADDRESS")
	viper.BindEnv("grpc.Addr", "CHAT_RPC_ADDRESS")
	viper.BindEnv("registryAddress", "CHAT_REGISTRY_ADDRESS")
	viper.BindEnv("tracerAddress", "CHAT_TRACER_ADDRESS")
	viper.BindEnv("debug", "CHAT_DEBUG")
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
		trace.WithServiceName(constvar.ServiceConnect),
		trace.WithEndpoint(conf.Conf.TracerAddress))
	if err != nil {
		log.Fatalf("Failed to init tracer: %v", err)
	}
}

//run 连接层服务启动
func run() {

	grpcServer := server.NewGRPCServer(&conf.Conf.GRPC)
	app := connect.New(
		connect.WithRegistry(
			consul.NewRegistry(
				registry.Addrs(conf.Conf.RegistryAddress),
				consul.TCPCheck(time.Second*30),
			),
		),
		connect.WithTransport(grpcServer),
	)

	//注册处理器
	pb.RegisterConnectServer(grpcServer, handler.New(app))

	//初始化
	app.Init(connect.WithServer(server.NewWsServer(app, conf.Conf.WsAddress)))

	//start server
	if err := app.Run(); err != nil {
		log.Fatalf("App Run err: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Connect Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Connect Server is exiting")
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
