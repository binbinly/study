package conf

import (
	"log"

	"chat/internal/conf"
	"chat/pkg/connect/tcp"
	"chat/pkg/connect/ws"
	logger "chat/pkg/log"
	"chat/pkg/net/grpc"
	"chat/pkg/registry"
	"chat/pkg/trace"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	App        AppConfig
	TCP        tcp.Config
	Ws         ws.Config
	GrpcClient grpc.ClientConfig
	GrpcServer grpc.ServerConfig
	Registry   registry.Config
	Logger     logger.Config
	Trace      trace.Config
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

// AppConfig app config
type AppConfig struct {
	Name     string
	Host     string
	ServerID string
	Env      string
	Debug    bool
}
