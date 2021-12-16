package app

import "common/wrap"

//Option 选项
type Option func(o *Options)

//Options 服务选项
type Options struct {
	client   bool
	name     string
	version  string
	migrate  func()
	authFunc wrap.AuthFunc
}

//WithMigrate 设置数据库迁移
func WithMigrate(migrate func()) Option {
	return func(o *Options) {
		o.migrate = migrate
	}
}

//WithAuthFunc 设置身份验证
func WithAuthFunc(authFunc wrap.AuthFunc) Option {
	return func(o *Options) {
		o.authFunc = authFunc
	}
}

//WithName with
func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

//WithVersion with
func WithVersion(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}

//WithClient with
func WithClient() Option {
	return func(o *Options) {
		o.client = true
	}
}