package conf

import (
	"log"

	"chat/internal/conf"
	logger "chat/pkg/log"
	"chat/pkg/net/tracing"
	"chat/pkg/registry"
	"chat/pkg/server/grpc"
	"chat/pkg/server/tcp"
	"chat/pkg/server/ws"
)

var Conf = &Config{}

// Config global config
type Config struct {
	Name       string
	Host       string
	ServerId   string
	LogicName  string
	Tcp        tcp.Config
	Ws         ws.Config
	GrpcClient grpc.ClientConfig
	GrpcServer grpc.ServerConfig
	Registry   registry.Config
	Logger     logger.Config
	Jaeger     tracing.Config
	Prometheus conf.PrometheusConfig
	Sentry     conf.SentryConfig
}

// Init init config
func Init(cfg string) {
	v, err := conf.LoadConfig(cfg)
	if err != nil {
		log.Fatalf("load config err:%v", err)
	}
	defaultConf(v)
	Conf = new(Config)
	err = v.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("init config err:%v", err)
	}
	conf.WatchConfig(v)
}
