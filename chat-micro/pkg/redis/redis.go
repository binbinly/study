package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

//Client redis 客户端
var Client *redis.Client

// Nil redis 返回为空
const Nil = redis.Nil

// Success redis成功标识
const Success = 1

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
}

// Init 实例化一个redis client
func Init(c *Config) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
		PoolSize:     c.PoolSize,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("[redis] redis ping err: %+v", err)
	}
	// 追踪
	Client.AddHook(NewTracingHook())

	return Client
}

//DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		MinIdleConn:  2,
		DialTimeout:  5,
		ReadTimeout:  3,
		WriteTimeout: 3,
		PoolSize:     5,
		PoolTimeout:  300,
	}
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:         "192.168.8.76:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     5,
		PoolTimeout:  300,
	})
	return
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

// Close 关闭连接
func Close() error {
	if Client == nil {
		return nil
	}
	return Client.Close()
}
