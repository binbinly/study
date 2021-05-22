package registry

import "time"

type Options struct {
	HeartBeat    int64         //心跳
	Timeout      time.Duration //超时
	Addr         []string      //服务地址
	RegistryPath string        //注册路径
}

type Option func(opts *Options)

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithAddr(addr []string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}

func WithRegistryPath(path string) Option {
	return func(opts *Options) {
		opts.RegistryPath = path
	}
}

func WithHeartBeat(heartHeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartHeat
	}
}
