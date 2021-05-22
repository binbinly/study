package conf

import (
	"log"

	"chat/internal/conf"
	logger "chat/pkg/log"
	"chat/pkg/net/tracing"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
	"chat/pkg/server/grpc"
)

var Conf = &Config{}

// Config global config
type Config struct {
	Consul     string
	Connect    Connect
	Queue      iqueue.Config
	Redis      redis.Config
	GrpcClient grpc.ClientConfig
	Logger     logger.Config
	Jaeger     tracing.Config
	Prometheus conf.PrometheusConfig
	Sentry     conf.SentryConfig
}

// Connect
type Connect struct {
	RoutineChan int
	RoutineSize int
	ServiceName string
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
