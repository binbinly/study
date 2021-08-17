package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"chat/app/center"
	"chat/app/center/conf"
	"chat/app/center/server"
	"chat/internal/orm"
	logger "chat/pkg/log"
	"chat/pkg/redis"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	"chat/pkg/trace"
)

func init() {
	centerCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/center/center.yaml)")
}

var centerCmd = &cobra.Command{
	Use:   "center",
	Short: "chat center server start",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/center/center.yaml"
		}
		conf.Init(cfg)
		centerStart()
	},
}

//centerStart 中心服务启动
func centerStart() {
	// init log
	logger.InitLog(&conf.Conf.Logger)
	// register consul plugin
	registry.RegisterPlugin(consul.NewConsul())
	// init registry
	rs, err := registry.InitRegistry(context.Background(), conf.Conf.Registry.Name,
		registry.WithAddr([]string{conf.Conf.Registry.Host}),
	)
	if err != nil {
		log.Fatalf("failed to init register: %v", err)
	}
	// init tracer
	if conf.Conf.Trace.Enable {
		_, err = trace.Init(conf.Conf.App.Name, conf.Conf.Trace.GetTraceConfig())
		if err != nil {
			log.Fatalf("failed to init trace: %v", err)
		}
	}
	// init db
	orm.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// init service
	svc := center.New(conf.Conf)
	// init grpc server
	grpcSrv := server.NewGRPCServer(conf.Conf, rs, svc)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Center Server is exiting") // close http server
			// close grpc server
			grpcSrv.Stop()
			// close service
			if err = svc.Close(); err != nil {
				log.Printf("service close err:%v\n", err)
			}
			// close redis
			if err = redis.Close(); err != nil {
				log.Printf("redis close err:%v\n", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
