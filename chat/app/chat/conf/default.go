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
		"Name":        "chat_logic",
		"Host":        localIP,
		"ServerID":    strings.ReplaceAll(localIP, ".", ""),
		"Mode":        "debug",
		"JwtSecret":   "Your-Jwt-Secret",
		"JwtTimeout":  86400,
		"Debug":       true,
		"Env":         app.EnvDev,
		"MaxLimit":    1000,
		"IPLimit":     100,
		"IPLimitExpr": "5m",
	})
	v.SetDefault("http", map[string]interface{}{
		"Port":         9050,
		"ReadTimeout":  "5s",
		"WriteTimeout": "5s",
	})
	v.SetDefault("mysql", map[string]interface{}{
		"Name":            "chat",
		"Addr":            "127.0.0.1:3306",
		"UserName":        "root",
		"Password":        "123456",
		"TablePrefix":     "",
		"Debug":           true,
		"MaxIdleConn":     10,
		"MaxOpenConn":     100,
		"ConnMaxLifeTime": "60m",
	})
	v.SetDefault("redis", map[string]interface{}{
		"Addr":         "127.0.0.1:6379",
		"Password":     "",
		"DB":           0,
		"MinIdleConn":  30,
		"DialTimeout":  "60s",
		"ReadTimeout":  "500ms",
		"WriteTimeout": "500ms",
		"PoolSize":     500,
		"PoolTimeout":  240,
	})
	v.SetDefault("elastic", map[string]interface{}{
		"Host": "http://127.0.0.1:9200",
	})
	v.SetDefault("queue", map[string]interface{}{
		"Plugin":  "redis",
		"Channel": "message",
		"Nsq": map[string]interface{}{
			"ProdHost": "127.0.0.1:4150",
			"Topic":    "message",
			"Channel":  "message",
		},
		"Rabbitmq": map[string]interface{}{
			"Addr":      "guest:guest@localhost:5672/",
			"QueueName": "message",
		},
	})
	v.SetDefault("grpcServer", map[string]interface{}{
		"Network":           "tcp",
		"Port":              20007,
		"Timeout":           "5s",
		"QPSLimit":          100,
		"IdleTimeout":       "15s",
		"MaxLifeTime":       "30s",
		"ForceCloseWait":    "5s",
		"KeepAliveInterval": "5s",
		"KeepAliveTimeout":  "1s",
	})
	v.SetDefault("grpcClient", map[string]interface{}{
		"ServiceName":      "center",
		"QPSLimit":         100,
		"Timeout":          "5s",
		"KeepAliveTime":    "15s",
		"KeepAliveTimeout": "1s",
	})
	v.SetDefault("registry", map[string]interface{}{
		"Name": "consul",
		"Host": "127.0.0.1:8500",
	})
}
