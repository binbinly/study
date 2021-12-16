package constvar

const (
	//ExchangeWarehouse 仓储服务rabbitmq交换器名
	ExchangeWarehouse = "stock.event.exchange"
	//ExchangeOrder 订单服务rabbitmq交换器名
	ExchangeOrder = "order.event.exchange"
)

const (
	//TopicWarehouse 仓储服务任务名
	TopicWarehouse = "mall.warehouse.stock"
	//TopicOrderSeckill 订单服务秒杀任务名
	TopicOrderSeckill = "mall.order.seckill"
	//TopicTask 任务服队列
	TopicTask = "mall.task"
)

const (
	//KeyOrderCreate 订单创建路由键
	KeyOrderCreate  = "order.create.order"
	//KeyOrderRelease 订单过期路由键
	KeyOrderRelease = "order.release.order"
)