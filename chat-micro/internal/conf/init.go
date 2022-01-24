package conf

import (
	"errors"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// LoadConfig load config file from given path
func LoadConfig(cfg string) error {
	viper.SetConfigFile(cfg)    // 如果指定了配置文件，则解析指定的配置文件
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量
	viper.SetEnvPrefix("CHAT")  // 读取环境变量的前缀为 chat
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//配置文件未找到错误；如果需要可以忽略
			return errors.New("config file not found")
		}
		// 配置文件被找到，但产生了另外的错误
		return err
	}

	log.Println("Using config file:", viper.ConfigFileUsed())

	return nil
}

//WatchConfig 监控配置文件变化并热加载程序
func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
