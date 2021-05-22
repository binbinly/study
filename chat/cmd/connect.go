package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"chat/app/connect"
	"chat/app/connect/conf"
	"chat/app/connect/grpc"
	logger "chat/pkg/log"
	"chat/pkg/registry"
	"chat/pkg/registry/consul"
)

func init() {
	connectCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/connect.local.yaml)")
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "chat connect server start",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/connect.local.yaml"
		}
		conf.Init(cfg)
		connectStart()
	},
}

func connectStart() {
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
	// init service
	svc := connect.NewServer(conf.Conf, rs)
	svc.StartWsServer()
	// init grpc server
	grpcSrv := grpc.New(conf.Conf, svc, rs)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Connect Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Connect Server is exiting")
			if err := svc.Close(); err != nil {
				log.Printf("Server close err: %s", err)
			}
			//注销服务
			rs.Unregister(context.Background(), &registry.Service{Id:"h-"+conf.Conf.ServerId})
			rs.Unregister(context.Background(), &registry.Service{Id:conf.Conf.ServerId})
			grpcSrv.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
