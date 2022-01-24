package logic

import (
	"context"
	"time"

	"chat-micro/pkg/registry"
	"chat-micro/pkg/transport"
)

var (
	registerTTL      = time.Minute
	registerInterval = time.Second * 30
)

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	ctx context.Context

	registry         registry.Registry
	registerTTL      time.Duration
	registerInterval time.Duration
	servers          []transport.Server
}

// WithID with app id
func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// WithVersion with a version
func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// WithServer with a server
func WithServer(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}

// WithRegistry with service registry.
func WithRegistry(r registry.Registry) Option {
	return func(o *options) {
		o.registry = r
	}
}

// WithRegistryTTL with service registryTTL.
func WithRegistryTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.registerTTL = ttl
	}
}

// WithRegistryInterval with service registryInterval.
func WithRegistryInterval(interval time.Duration) Option {
	return func(o *options) {
		o.registerInterval = interval
	}
}
