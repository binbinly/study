package conf

import (
	"github.com/spf13/viper"

	"chat/internal/conf"
)

func defaultConf(v *viper.Viper) {
	conf.DefaultConf(v)
	v.SetDefault("app", map[string]interface{}{
		"Name":        "chat_task",
		"RoutineChan": 1024,
		"RoutineSize": 8,
	})
	v.SetDefault("consul", "127.0.0.1:8500")
	v.SetDefault("queue", map[string]interface{}{
		"Plugin":  "redis",
		"Channel": "message",
		"Nsq": map[string]interface{}{
			"ConsumerHost": []string{"127.0.0.1:4161"},
			"Topic":        "message",
			"Channel":      "message",
			"MaxAttempts":  3,
		},
		"Rabbitmq": map[string]interface{}{
			"Addr":      "guest:guest@localhost:5672/",
			"QueueName": "message",
		},
	})
	v.SetDefault("grpcClient", map[string]interface{}{
		"ServiceName":      "chat_connect",
		"Timeout":          "5s",
		"QPSLimit":         100,
		"KeepAliveTime":    "15s",
		"KeepAliveTimeout": "1s",
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
}
