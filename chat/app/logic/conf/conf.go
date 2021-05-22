package conf

import (
	"log"
	"time"

	"chat/internal/conf"
	"chat/pkg/database/orm"
	logger "chat/pkg/log"
	"chat/pkg/net/tracing"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
	"chat/pkg/registry"
	"chat/pkg/server/grpc"
)

var Conf = &Config{}

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

// Config global config
type Config struct {
	App        AppConfig
	Http       HttpConfig
	Sms        SmsConfig
	MySQL      orm.Config
	Redis      redis.Config
	GrpcServer grpc.ServerConfig
	GrpcClient grpc.ClientConfig
	Queue      iqueue.Config
	Registry   registry.Config
	Logger     logger.Config
	Jaeger     tracing.Config
	Prometheus conf.PrometheusConfig
	Sentry     conf.SentryConfig
}

// AppConfig app config
type AppConfig struct {
	Name       string
	Host       string
	ServerId   string
	LogicName  string
	Mode       string
	PprofPort  string
	JwtSecret  string
	JwtTimeout int64
	Debug      bool
}

type HttpConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// LimitConfig 限流器
type LimitConfig struct {
	Enable bool
	Qps    int
}

type SmsConfig struct {
	IsReal bool `yaml:"is_real"` //是否真实发送
}
