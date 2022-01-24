package conf

import (
	"github.com/spf13/viper"
	"log"

	"chat-micro/internal/conf"
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
	RoutineNum      int
	RoutineSize     int
	Debug           bool
	Redis           redis.Config
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
	viper.SetDefault("brokerAddress", "nats://127.0.0.1:4222")
	viper.SetDefault("routineNum", 4)
	viper.SetDefault("routineSize", 128)
	viper.SetDefault("redis", map[string]interface{}{
		"Addr":         "127.0.0.1:6379",
		"Password":     "",
		"DB":           0,
		"MinIdleConn":  30,
		"DialTimeout":  60,
		"ReadTimeout":  3,
		"WriteTimeout": 3,
		"PoolSize":     500,
		"PoolTimeout":  240,
	})
}
