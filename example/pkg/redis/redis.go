package redis

import (
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"log"
)

//Client redis 客户端
var Client *redis.Client

const (
	// ErrRedisNotFound not exist in redis
	ErrRedisNotFound = redis.Nil
	// DefaultRedisName default redis name
	DefaultRedisName = "default"
)

// Config redis config
type Config struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConn  int
	DialTimeout  int64
	ReadTimeout  int64
	WriteTimeout int64
	PoolSize     int
	PoolTimeout  int64
	EnableTrace  bool
}

// Init init a default redis instance
func Init(c *Config) *redis.Client {
	clientManager := NewRedisManager()
	rdb, err := clientManager.GetClient(DefaultRedisName, c)
	if err != nil {
		log.Fatalf("init redis err: %s", err)
	}
	Client = rdb

	return rdb
}

// InitCustomTestRedis 实例化自定义的测试客户端
func InitCustomTestRedis(addr, password string) {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
	})
	fmt.Println("customize redis addr:", addr)
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	Client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
