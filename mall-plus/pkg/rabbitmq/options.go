package rabbitmq

import "github.com/streadway/amqp"

//Option 选项执行
type Option func(o *Options)

//Options 队列选项
type Options struct {
	name         string // 名称
	durable      bool   // 是否持久化
	autoDelete   bool   // 是否自动删除，在最后一个consumer断开连接后，删除
	exclusive    bool   // 是否独占队列，只能被当前连接访问，连接关闭队列也会删除
	noWait       bool
	args         amqp.Table
}

//WithName 设置名称
func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

//WithDurable 设置
func WithDurable() Option {
	return func(o *Options) {
		o.durable = true
	}
}

//WithAutoDelete 设置
func WithAutoDelete() Option {
	return func(o *Options) {
		o.autoDelete = true
	}
}

//WithExclusive 设置
func WithExclusive() Option {
	return func(o *Options) {
		o.exclusive = true
	}
}

//WithNoWait 设置
func WithNoWait() Option {
	return func(o *Options) {
		o.noWait = true
	}
}

//WithArgs 设置
func WithArgs(args amqp.Table) Option {
	return func(o *Options) {
		o.args = args
	}
}
