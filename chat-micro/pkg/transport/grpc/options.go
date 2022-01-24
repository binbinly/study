package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// Option is gRPC server option.
type Option func(o *options)

type options struct {
	network    string
	address    string
	maxMsgSize int
	keepalive  bool
	timeout    time.Duration
	inters     []grpc.UnaryServerInterceptor
	grpcOpts   []grpc.ServerOption
}

// Network with server network.
func Network(network string) Option {
	return func(s *options) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) Option {
	return func(s *options) {
		s.address = addr
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) Option {
	return func(s *options) {
		s.timeout = timeout
	}
}
// MaxMsgSize with server maxMsgSize.
func MaxMsgSize(size int) Option {
	return func(s *options) {
		s.maxMsgSize = size
	}
}

// Keepalive with server keepalive.
func Keepalive() Option {
	return func(s *options) {
		s.keepalive = true
	}
}

// UnaryInterceptor returns a ServerOption that sets the UnaryServerInterceptor for the server.
func UnaryInterceptor(in ...grpc.UnaryServerInterceptor) Option {
	return func(s *options) {
		s.inters = in
	}
}

// Options with grpc options.
func Options(opts ...grpc.ServerOption) Option {
	return func(s *options) {
		s.grpcOpts = opts
	}
}

func newOptions(opt ...Option) options {
	opts := options{
		network:    "tcp",
		address:    ":0",
		timeout:    5 * time.Second,
		maxMsgSize: 1024 * 1024 * 4,
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
