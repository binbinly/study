package redis

import (
	"time"

	"github.com/go-redis/redis"

	"chat/pkg/log"
)

// RedisClient redis 客户端
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

	_, err := Client.Ping().Result()
	if err != nil {
		log.Panicf("[redis] redis ping err: %+v", err)
	}
	return Client
}

// Close 关闭连接
func Close() error {
	if Client == nil {
		return nil
	}
	return Client.Close()
}
