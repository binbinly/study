package conf

import (
	"log"

	"chat/internal/conf"
	"chat/pkg/database/orm"
	logger "chat/pkg/log"
	"chat/pkg/net/grpc"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
	"chat/pkg/registry"
	"chat/pkg/trace"
)

//Conf 全局配置
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
	Sms        SmsConfig
	MySQL      orm.Config
	Redis      redis.Config
	GrpcServer grpc.ServerConfig
	Queue      iqueue.Config
	Registry   registry.Config
	Logger     logger.Config
	Trace      trace.Config
}

// AppConfig app config
type AppConfig struct {
	Name       string
	Host       string
	ServerID   string
	JwtSecret  string
	JwtTimeout int64
}

//SmsConfig 短信配置
type SmsConfig struct {
	IsReal bool `yaml:"is_real"` //是否真实发送
}
