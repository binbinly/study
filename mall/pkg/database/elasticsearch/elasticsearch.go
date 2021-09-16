package elasticsearch

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

//Client 全局 elastic客户端
var Client *elastic.Client

//Config elastic配置项
type Config struct {
	Sniff    bool
	Host     string
	Username string
	Password string
}

//NewClient 实例化elastic客户端
func NewClient(c *Config) *elastic.Client {
	var err error
	Client, err = elastic.NewClient(elastic.SetSniff(c.Sniff), elastic.SetURL(c.Host),
		elastic.SetBasicAuth(c.Username, c.Password))
	if err != nil {
		log.Fatal("new elastic client err", err)
	}
	info, code, err := Client.Ping(c.Host).Do(context.Background())
	if err != nil {
		log.Fatal("ping elastic err", err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return Client
}
