package conf

import "pkg/redis"

//Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	Redis redis.Config
	Sms   SmsConfig `yaml:"sms"`
	Eth   EthConfig `yaml:"eth"`
}

// SmsConfig 短信配置
type SmsConfig struct {
	IsReal bool
}

// EthConfig 以太坊配置
type EthConfig struct {
	NetworkID  int
	NetworkUrl string
}
