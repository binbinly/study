package conf

import (
	"github.com/spf13/viper"
)

func defaultConf() {
	viper.SetDefault("registryAddress", "127.0.0.1:8500")
	viper.SetDefault("brokerAddress", "nats://127.0.0.1:4222")
	viper.SetDefault("JwtTimeout", 86400)
	viper.SetDefault("http", map[string]interface{}{
		"Addr":         ":9050",
		"ReadTimeout":  "5s",
		"WriteTimeout": "5s",
	})
	viper.SetDefault("grpc", map[string]interface{}{
		"Network":    "tcp",
		"Addr":       ":20007",
		"Timeout":    "5s",
		"MaxMsgSize": 4194304,
	})
	viper.SetDefault("mysql", map[string]interface{}{
		"Name":            "chat",
		"Addr":            "127.0.0.1:3306",
		"UserName":        "root",
		"Password":        "123456",
		"TablePrefix":     "",
		"Debug":           true,
		"MaxIdleConn":     10,
		"MaxOpenConn":     100,
		"ConnMaxLifeTime": 60,
	})
	viper.SetDefault("redis", map[string]interface{}{
		"Addr":         "127.0.0.1:6379",
		"Password":     "",
		"DB":           0,
		"MinIdleConn":  30,
		"DialTimeout":  60,
		"ReadTimeout":  3,
		"WriteTimeout": 3,
		"PoolSize":     500,
		"PoolTimeout":  240,
	})
	viper.SetDefault("minio", map[string]interface{}{
		"Endpoint":     "127.0.0.1:9000",
		"AccessID":     "minioadmin",
		"SecretAccess": "minioadmin",
		"Bucket":       "group1",
		"Region":       "us-east-1",
	})
}
