package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"chat/app/task"
	"chat/app/task/conf"
	logger "chat/pkg/log"
	"chat/pkg/redis"
)

func init() {
	taskCmd.Flags().StringVarP(&cfg, "config", "c", "", "config file (default is $ROOT/config/task.local.yaml)")
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "chat task server start",
	Run: func(cmd *cobra.Command, args []string) {
		if cfg == "" {
			cfg = "./config/task.local.yaml"
		}
		conf.Init(cfg)
		taskStart()
	},
}

func taskStart() {
	// init log
	logger.InitLog(&conf.Conf.Logger)
	// init redis
	redis.Init(&conf.Conf.Redis)
	// new task and start
	t := task.New(conf.Conf)
	t.Start()

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("chat task get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("chat task server exit")
			t.Close()
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
