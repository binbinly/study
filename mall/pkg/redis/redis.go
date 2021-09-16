package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"

	"mall/pkg/log"
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
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

// Init 实例化一个redis client
func Init(c *Config) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("[redis] redis ping err: %+v", err)
	}
	// 追踪
	Client.AddHook(NewTracingHook())

	return Client
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

// Close 关闭连接
func Close() error {
	if Client == nil {
		return nil
	}
	return Client.Close()
}
