package conf

import (
	"common/conf"
	"pkg/database/mysql"
	"pkg/redis"
)

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	DFS   conf.DFSConfig
	MySQL mysql.Config
	Redis redis.Config
}
