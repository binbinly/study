package conf

import (
	"log"

	"chat/internal/conf"
	logger "chat/pkg/log"
	"chat/pkg/net/grpc"
	"chat/pkg/queue/iqueue"
	"chat/pkg/redis"
	"chat/pkg/trace"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	Consul     string
	App        AppConfig
	Queue      iqueue.Config
	Redis      redis.Config
	GrpcClient grpc.ClientConfig
	Logger     logger.Config
	Trace      trace.Config
}

// AppConfig app配置
type AppConfig struct {
	RoutineChan int
	RoutineSize int
	Name        string
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
