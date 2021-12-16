package conf

import (
	"pkg/database/mysql"
	"pkg/redis"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	App   AppConfig
	MySQL mysql.Config
	Redis redis.Config
}

// AppConfig app config
type AppConfig struct {
	Name       string
	Debug      bool
}
