package conf

import (
	"log"
	"time"

	"github.com/spf13/viper"

	"chat-micro/internal/conf"
	"chat-micro/pkg/database/mysql"
	"chat-micro/pkg/minio"
	"chat-micro/pkg/redis"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	Name            string
	BrokerAddress   string
	RegistryAddress string
	TracerAddress   string
	JwtSecret       string
	JwtTimeout      int64
	Debug           bool
	MySQL           mysql.Config
	Redis           redis.Config
	Minio           minio.Config
	HTTP            HTTPConfig
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

// HTTPConfig server config.
type HTTPConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// GRPCConfig server config.
type GRPCConfig struct {
	Network    string
	Addr       string
	Timeout    time.Duration
	MaxMsgSize int
}