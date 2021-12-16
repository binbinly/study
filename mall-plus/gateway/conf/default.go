package conf

import (
	"github.com/spf13/viper"
)

func defaultConf(v *viper.Viper) {
	v.SetDefault("app", map[string]interface{}{
		"Name":      "gateway",
		"Version":   "latest",
		"JwtSecret": "Your-Jwt-Secret",
		"Debug":     true,
	})
	v.SetDefault("http", map[string]interface{}{
		"Addr":         ":9520",
		"ReadTimeout":  "5s",
		"WriteTimeout": "5s",
	})
	v.SetDefault("registry", map[string]interface{}{
		"Name": "consul",
		"Host": "127.0.0.1:8500",
	})
	v.SetDefault("services", map[string]interface{}{
		"Cart": map[string]interface{}{
			"Name":     "mall.cart",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
		"Market": map[string]interface{}{
			"Name":     "mall.market",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
		"Member": map[string]interface{}{
			"Name":     "mall.member",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
		"Order": map[string]interface{}{
			"Name":     "mall.order",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
		"Product": map[string]interface{}{
			"Name":     "mall.product",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
		"Seckill": map[string]interface{}{
			"Name":     "mall.seckill",
			"Timeout":  "5s",
			"QPSLimit": "500",
		},
	})
}
