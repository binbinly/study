package task

import "chat-micro/pkg/broker"

const (
	_version         = "latest"
	_routineSIze     = 128
	_routineNum      = 4
	_registryAddress = "127.0.0.1:8500"
)

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	name     string
	version  string
	metadata map[string]string

	topic           string
	registryAddress string
	connectName     string
	routineSize     int //任务大小
	routineNum      int //任务个数

	broker     broker.Broker
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

// WithRegistryAddress with a registryAddress
func WithRegistryAddress(addr string) Option {
	return func(o *options) {
		o.registryAddress = addr
	}
}

// WithConnectName with a connectName
func WithConnectName(connectName string) Option {
	return func(o *options) {
		o.connectName = connectName
	}
}

// WithRoutineSize with a routineSize
func WithRoutineSize(routineSize int) Option {
	return func(o *options) {
		o.routineSize = routineSize
	}
}

// WithRoutineNum with a routineNum
func WithRoutineNum(routineNum int) Option {
	return func(o *options) {
		o.routineNum = routineNum
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

func newOptions(opt ...Option) options {
	opts := options{
		version:         _version,
		registryAddress: _registryAddress,
		routineSize:     _routineSIze,
		routineNum:      _routineNum,
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
