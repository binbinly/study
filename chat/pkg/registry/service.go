package registry

// 服务抽象
type Service struct {
	Id    string
	Name  string
	IP    string
	Port  int
	Check Check
}

// 健康检查
type Check struct {
	GRPC string // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	HTTP string // http 支持，执行健康检查的地址
	TCP  string	// tcp 支持,执行健康检查的地址
}
