package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"gateway/app"
	"gateway/conf"
	"gateway/internal/interceptor/sentinel"
	"gateway/internal/router"
	"gateway/internal/server"
	"pkg/logger"
	"pkg/registry"
	"pkg/registry/consul"
	"pkg/trace"
)

var (
	configYml string
	//StartCmd server cmd
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API Gateway",
		Example:      "gateway server -c default.yml",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "default.yaml", "Start server with provided configuration file")
}

func setup() {
	conf.Init(configYml)
	// init log
	l, err := logger.NewZapLogger()
	if err != nil {
		log.Fatalf("Failed to init zap log err: %v", err)
	}
	logger.DefaultLogger = l

	// init tracer
	_, err = trace.InitTracerProvider(
		trace.WithServiceName(conf.Conf.App.Name),
		trace.WithEndpoint(conf.Conf.Trace.Endpoint))
	if err != nil {
		log.Fatalf("Failed to init tracer: %v", err)
	}
	// init sentinel
	sentinel.Init(conf.Conf.App.Name)
}

func run() {
	mux := router.NewServeMux()

	a := app.New(
		app.WithName(conf.Conf.App.Name),
		app.WithVersion(conf.Conf.App.Version),
		app.WithRegistry(consul.NewRegistry(registry.Addrs(conf.Conf.Registry.Addr),
			consul.TCPCheck(time.Second*30))),
		app.WithServices(&conf.Conf.Services),
		app.WithMux(mux),
		app.WithServer(
			server.NewHTTPServer(&conf.Conf.HTTP, mux)))

	//register service handler
	a.RegisterAll()
	//start server
	if err := a.Run(); err != nil {
		log.Fatalf("App Run err: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-a.Ctx().Done():
			log.Printf("Ctx done err: %v", a.Ctx().Err())
			return
		case s := <-quit:
			log.Printf("Server receive a quit signal: %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				log.Println("Server is exiting") // close http server
				if err := a.Stop(); err != nil {
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
}
