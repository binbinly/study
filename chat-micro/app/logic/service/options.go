package service

import (
	"github.com/google/uuid"

	"chat-micro/pkg/broker"
	"chat-micro/pkg/minio"
)

const (
	_version    = "latest"
	_jwtTimeout = 86400
)

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	topic      string
	jwtSecret  string
	jwtTimeout int64
	smsReal    bool

	broker  broker.Broker
	storage *minio.Storage
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

// WithSmsReal with a smsReal
func WithSmsReal() Option {
	return func(o *options) {
		o.smsReal = true
	}
}

// WithTopic with a server
func WithTopic(topic string) Option {
	return func(o *options) {
		o.topic = topic
	}
}

// WithBroker with a broker
func WithBroker(b broker.Broker) Option {
	return func(o *options) {
		o.broker = b
	}
}

// WithJwtSecret with a jwtSecret
func WithJwtSecret(jwtSecret string) Option {
	return func(o *options) {
		o.jwtSecret = jwtSecret
	}
}

// WithStorage with a storage
func WithStorage(s *minio.Storage) Option {
	return func(o *options) {
		o.storage = s
	}
}

func newOptions(opt ...Option) options {
	opts := options{
		version:    _version,
		jwtTimeout: _jwtTimeout,
	}

	if id, err := uuid.NewUUID(); err == nil {
		opts.id = id.String()
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
