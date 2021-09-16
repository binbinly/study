package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"mall/app/conf"
	"mall/app/model"
	"mall/app/routers"
	"mall/app/server"
	"mall/app/service"
	"mall/pkg/database/orm"
	logger "mall/pkg/log"
	"mall/pkg/redis"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "mall server -c config/default.yml",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/default.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
}

func setup() {
	conf.Init(configYml)
	// init log
	logger.InitLog(&conf.Conf.Logger)
	// init db
	model.Init(&conf.Conf.MySQL)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// Set gin mode.
	gin.SetMode(conf.Conf.App.Mode)
}

func run() {
	// init http server
	httpSrv := server.NewHTTPServer(conf.Conf, routers.NewRouter(&conf.Conf.App))
	// init service
	svc := service.New(conf.Conf)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-quit
		log.Printf("Server receive a quit signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Server is exiting") // close http server
			if err := httpSrv.Stop(); err != nil {
				log.Printf("Server shutdown err: %s", err)
			}
			// close service
			if err := svc.Close(); err != nil {
				log.Printf("service close err:%v", err)
			}
			// close db
			if err := orm.CloseDB(); err != nil {
				log.Printf("mysql close err:%v", err)
			}
			// close redis
			if err := redis.Close(); err != nil {
				log.Printf("redis close err:%v", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
