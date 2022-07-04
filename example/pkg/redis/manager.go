package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

// Manager define a redis manager
type Manager struct {
	clients map[string]*redis.Client
	*sync.RWMutex
}

// NewRedisManager create a redis manager
func NewRedisManager() *Manager {
	return &Manager{
		clients: make(map[string]*redis.Client),
		RWMutex: &sync.RWMutex{},
	}
}

// GetClient get a redis instance
func (r *Manager) GetClient(name string, c *Config) (*redis.Client, error) {
	// get client from map
	r.RLock()
	if client, ok := r.clients[name]; ok {
		r.RUnlock()
		return client, nil
	}
	r.RUnlock()

	// create a redis client
	r.Lock()
	defer r.Unlock()
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
		PoolSize:     c.PoolSize,
	})

	// check redis if is ok
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	// hook tracing (using open telemetry)
	if c.EnableTrace {
		rdb.AddHook(redisotel.NewTracingHook())
	}
	r.clients[name] = rdb

	return rdb, nil
}