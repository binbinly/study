package conf

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"mall/pkg/database/orm"
	logger "mall/pkg/log"
	"mall/pkg/net/http"
	"mall/pkg/redis"
)

//Conf 全局配置
var Conf = &Config{}

// Init init config
func Init(cfg string) {
	v, err := loadConfig(cfg)
	if err != nil {
		log.Fatalf("load config err:%v", err)
	}
	defaultConf(v)
	Conf = new(Config)
	err = v.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("init config err:%v", err)
	}
	watchConfig(v)
}

// loadConfig load config file from given path
func loadConfig(cfg string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(cfg)       // 如果指定了配置文件，则解析指定的配置文件
	v.SetConfigType("yaml")    // 设置配置文件格式为YAML
	v.AutomaticEnv()           // 读取匹配的环境变量
	viper.SetEnvPrefix("mall") // 读取环境变量的前缀为 mall
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//配置文件未找到错误；如果需要可以忽略
			return nil, errors.New("config file not found")
		}
		// 配置文件被找到，但产生了另外的错误
		return nil, err
	}
	log.Println("Using config file:", viper.ConfigFileUsed())

	return v, nil
}

//watchConfig 监控配置文件变化并热加载程序
func watchConfig(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

// Config global config
type Config struct {
	App    AppConfig
	HTTP   http.ServerConfig
	Eth    EthConfig
	MySQL  orm.Config
	Redis  redis.Config
	Logger logger.Config
}

// AppConfig app config
type AppConfig struct {
	Name        string
	DfsUrl      string
	Env         string
	Mode        string
	JwtSecret   string
	JwtTimeout  int64
	Debug       bool
	MaxLimit    int
	IPLimit     int
	IPLimitExpr time.Duration
}

// EthConfig 以太坊配置
type EthConfig struct {
	NetworkID  int
	NetworkUrl string
}
