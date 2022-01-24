package conf

import (
	"log"
	"time"

	"github.com/spf13/viper"

	"chat-micro/internal/conf"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	Name            string
	WsAddress       string
	TcpAddress      string
	RegistryAddress string
	TracerAddress   string
	Debug           bool
	GRPC            GRPCConfig
}

// Init init config
func Init(cfg string) {
	if err := conf.LoadConfig(cfg); err != nil {
		log.Fatalf("load config err:%v", err)
	}
	defaultConf()
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("init config err:%v", err)
	}
	log.Printf("config: %+v\n", Conf)
	conf.WatchConfig()
}

func defaultConf() {
	viper.SetDefault("registryAddress", "127.0.0.1:8500")
	viper.SetDefault("wsAddress", ":9070")
	viper.SetDefault("tcpAddress", ":9060")
	viper.SetDefault("grpc", map[string]interface{}{
		"Network":    "tcp",
		"Addr":       ":20005",
		"Timeout":    "5s",
		"MaxMsgSize": 4194304,
	})
}

// GRPCConfig server config.
type GRPCConfig struct {
	Network    string
	Addr       string
	Timeout    time.Duration
	MaxMsgSize int
}
