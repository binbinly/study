package conf

import (
	"strings"

	"github.com/spf13/viper"

	"chat/internal/conf"
	"chat/pkg/app"
	"chat/pkg/net/ip"
)

func defaultConf(v *viper.Viper) {
	conf.DefaultConf(v)
	localIP := ip.GetLocalIP()
	v.SetDefault("app", map[string]interface{}{
		"Name":     "chat_connect",
		"Host":     localIP,
		"ServerID": strings.ReplaceAll(localIP, ".", ""),
		"Debug":    true,
		"Env":      app.EnvDev,
	})
	v.SetDefault("tcp", map[string]interface{}{
		"Port":             9060,
		"MaxIpLimit":       0,
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
		"MaxIpLimit":       0,
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
		"ServiceName":      "center",
		"QPSLimit":         100,
		"Timeout":          "5s",
		"KeepAliveTime":    "15s",
		"KeepAliveTimeout": "1s",
	})
	v.SetDefault("grpcServer", map[string]interface{}{
		"Network":           "tcp",
		"Port":              20005,
		"QPSLimit":          100,
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
