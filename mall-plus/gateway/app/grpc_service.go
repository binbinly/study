package app

import "time"

//GRPCServices 网关代理的所有grpc服务
type GRPCServices struct {
	Cart    ServiceItem
	Market  ServiceItem
	Member  ServiceItem
	Order   ServiceItem
	Product ServiceItem
	Seckill ServiceItem
}

//ServiceItem service
type ServiceItem struct {
	Name     string
	Timeout  time.Duration
	QPSLimit float64
}
