package task

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"chat-micro/app/constvar"
	"chat-micro/app/task"
	"chat-micro/app/task/conf"
	"chat-micro/pkg/broker"
	"chat-micro/pkg/broker/nats"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/redis"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "task",
		Short:        "Start chat-micro task server",
		Example:      "chat-micro task -c config/task/default.yml",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/task/default.yml", "Start server with provided configuration file")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CHAT") // 读取环境变量的前缀为 chat
	viper.BindEnv("registryAddress", "CHAT_REGISTRY_ADDRESS")
	viper.BindEnv("tracerAddress", "CHAT_TRACER_ADDRESS")
	viper.BindEnv("brokerAddress", "CHAT_BROKER_ADDRESS")
	viper.BindEnv("redis.Addr", "CHAT_REDIS_ADDRESS")
	viper.BindEnv("redis.Password", "CHAT_REDIS_PWD")
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
}

//run 任务层启动
func run() {
	// init redis
	redis.Init(&conf.Conf.Redis)

	b := nats.NewBroker(broker.Addrs(conf.Conf.BrokerAddress))
	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	app := task.New(
		task.WithName(constvar.ServiceTask),
		task.WithConnectName(constvar.ServiceConnect),
		task.WithTopic(constvar.TaskTopic),
		task.WithBroker(b),
		task.WithRegistryAddress(conf.Conf.RegistryAddress))

	//start server
	app.Run()

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("chat-micro task get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("chat-micro task server exit")
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
