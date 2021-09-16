package conf

import (
	"github.com/spf13/viper"

	"mall/pkg/app"
)

func defaultConf(v *viper.Viper) {
	v.SetDefault("app", map[string]interface{}{
		"Name":        "chat_logic",
		"DfsUrl":      "http://127.0.0.1:9000",
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
		"Port":         9052,
		"ReadTimeout":  "5s",
		"WriteTimeout": "5s",
	})
	v.SetDefault("eth", map[string]interface{}{
		"NetworkID": 15,
		"NetworkUrl": "ws://localhost:8546",
	})
	v.SetDefault("mysql", map[string]interface{}{
		"Name":            "mall",
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
	v.SetDefault("logger", map[string]interface{}{
		"Development":     false,
		"DisableCaller":   false,
		"Encoding":        "json",
		"Level":           "INFO",
		"Name":            "mall",
		"Writers":         "file",
		"LoggerFile":      "./logs/mall.log",
		"LoggerWarnFile":  "./logs/mall.warn.log",
		"LoggerErrorFile": "./logs/mall.err.log",
	})
}
