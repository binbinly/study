package registry

//Options 选项结构
type Options struct {
	//Addr 服务地址
	Addr []string
}

//Option 选项
type Option func(opts *Options)

//WithAddr 设置地址
func WithAddr(addr []string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}
