package conf

import (
	"log"
	"time"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
)

//LoadFile 加载文件配置
func LoadFile(filepath string, path string, val interface{}) {
	load(file.NewSource(
		file.WithPath(filepath),
	), path, val)
}

//LoadConsul 加载consul配置
func LoadConsul(address string, path string, val interface{}) {
	if address == "" {
		log.Fatalln("consul address empty")
	}
	load(consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(address),
		//设置前缀，不设置默认为 /micro/config
		consul.WithPrefix("/mall/config"),
		//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
		source.WithEncoder(yaml.NewEncoder()),
	), path, val)
}

func load(s source.Source, path string, val interface{}) {
	// new yaml encoder
	enc := yaml.NewEncoder()

	// new config
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(enc),
			),
		),
	)

	// load the config from a file source
	if err := c.Load(s); err != nil {
		log.Fatalf("config load err: %v", err)
		return
	}

	go watch(c, path, val)
	//fmt.Println("data", c.Map())

	// read a database host
	if err := c.Get(path).Scan(val); err != nil {
		log.Fatalf("config scan product err: %v", err)
		return
	}
}

//watch 配置变化监听
func watch(c config.Config, path string, val interface{}) {
	for {
		w, err := c.Watch(path)
		if err != nil {
			logger.Warnf("[config] watch err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		v, err := w.Next()
		if err != nil {
			logger.Warnf("[config] next err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		if err = v.Scan(val); err != nil {
			logger.Warnf("[config] scan err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		logger.Info("[config] update conf: %v", val)
	}
}
