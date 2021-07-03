package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/app/logic/routers"
	"chat/app/logic/server"
	"chat/app/logic/service"
	"chat/pkg/database/orm"
	logger "chat/pkg/log"
	"chat/pkg/redis"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
	"chat/pkg/trace"
)

func init() {
	logicCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/logic.yaml)")
}

var logicCmd = &cobra.Command{
	Use:   "logic",
	Short: "chat logic server start",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/logic.yaml"
		}
		conf.Init(cfg)
		logicStart()
	},
}

//logicStart 逻辑业务服务启动
func logicStart() {
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
	model.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// Set gin mode.
	gin.SetMode(conf.Conf.App.Mode)
	// init http server
	httpSrv := server.NewHTTPServer(conf.Conf, rs, routers.NewRouter(&conf.Conf.App))
	// init service
	svc := service.New(conf.Conf)
	// init grpc server
	grpcSrv := server.NewGRPCServer(conf.Conf, rs, svc, routers.NewGrpcRouter())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Server is exiting") // close http server
			if err = httpSrv.Stop(); err != nil {
				log.Printf("Server shutdown err: %s", err)
			}
			// close grpc server
			grpcSrv.Stop()
			// close service
			if err = svc.Close(); err != nil {
				log.Printf("service close err:%v", err)
			}
			// close db
			if err = orm.CloseDB(); err != nil {
				log.Printf("mysql close err:%v", err)
			}
			// close redis
			if err = redis.Close(); err != nil {
				log.Printf("redis close err:%v", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
