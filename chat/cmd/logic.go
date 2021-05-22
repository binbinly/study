package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"chat/app/logic/conf"
	"chat/app/logic/grpc"
	"chat/app/logic/http"
	"chat/app/logic/model"
	"chat/app/logic/routers"
	"chat/app/logic/service"
	"chat/pkg/database/orm"
	logger "chat/pkg/log"
	"chat/pkg/redis"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
)

func init() {
	logicCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/logic.local.yaml)")
}

var logicCmd = &cobra.Command{
	Use:   "logic",
	Short: "chat logic server start",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/logic.local.yaml"
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
	// init db
	model.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// Set gin mode.
	gin.SetMode(conf.Conf.App.Mode)
	// init http server
	httpSrv := http.StartServer(rs)
	// init service
	svc := service.New(conf.Conf)
	// init grpc server
	grpcSrv := grpc.New(conf.Conf, rs, svc, routers.NewRouter())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Server is exiting") // close http server
			if err = httpSrv.Shutdown(ctx); err != nil {
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
			rs.Unregister(context.Background(), &registry.Service{Id: conf.Conf.App.ServerId})
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
