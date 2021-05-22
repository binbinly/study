package nsq

import (
	"github.com/nsqio/go-nsq"
)

// Config nsq消息队列
type Config struct {
	ProdHost     string
	ConsumerHost []string
	Topic        string
	Channel      string
	MaxInFlight  int    //并发处理消息数量 default 1
	MaxAttempts  uint16 //消息重试次数 default 5
}

func setting(conf *Config) *nsq.Config {
	c := nsq.NewConfig()
	if conf.MaxInFlight > 0 {
		c.MaxInFlight = conf.MaxInFlight
	}
	if conf.MaxAttempts > 0 {
		c.MaxAttempts = conf.MaxAttempts
	}
	return c
}
