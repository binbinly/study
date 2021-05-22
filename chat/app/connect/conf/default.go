package conf

import (
	"log"

	"github.com/spf13/viper"

	"chat/internal/conf"
	"chat/pkg/net/ip"
	"chat/pkg/utils"
)

func defaultConf(v *viper.Viper) {
	conf.DefaultConf(v)
	id, err := utils.GenShortID()
	if err != nil {
		log.Panicf("gen short id err:%v", err)
	}
	v.SetDefault("host", ip.GetLocalIP())
	v.SetDefault("name", "logic")
	v.SetDefault("serverId", id)
	v.SetDefault("LogicName",  "logic")
	v.SetDefault("tcp", map[string]interface{}{
		"Port":             9060,
		"Keepalive":        false,
		"HandshakeTimeout": "5s",
		"SendBuf":          4096,
		"ReceiveBuf":       4096,
		"MaxPacketSize":    4096,
		"MaxConn":          40000,
		"WorkerPoolSize":   16,
		"MaxWorkerTaskLen": 256,
		"MaxMsgChanLen":    256,
		"BucketSize":       32,
	})
	v.SetDefault("ws", map[string]interface{}{
		"Port":             9070,
		"WriteWait":        "10s",
		"PongWait":         "60s",
		"PingPeriod":       "54s",
		"MaxPacketSize":    4096,
		"ReadBufferSize":   4096,
		"WriteBufferSize":  4096,
		"MaxConn":          40000,
		"WorkerPoolSize":   16,
		"MaxWorkerTaskLen": 256,
		"MaxMsgChanLen":    256,
		"BucketSize":       16,
	})
	v.SetDefault("grpcClient", map[string]interface{}{
		"Timeout":          "5s",
		"KeepAliveTime":    "15s",
		"KeepAliveTimeout": "1s",
	})
	v.SetDefault("grpcServer", map[string]interface{}{
		"Network":           "tcp",
		"Port":              20005,
		"Timeout":           "5s",
		"IdleTimeout":       "15s",
		"MaxLifeTime":       "30s",
		"ForceCloseWait":    "5s",
		"KeepAliveInterval": "5s",
		"KeepAliveTimeout":  "1s",
	})
	v.SetDefault("registry", map[string]interface{}{
		"Name": "consul",
		"Host": "127.0.0.1:8500",
	})
}
