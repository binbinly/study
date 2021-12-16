package conf

import (
	"common/conf"
	"pkg/redis"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	DFS   conf.DFSConfig
	Redis redis.Config
}
