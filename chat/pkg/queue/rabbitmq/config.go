package rabbitmq

type Config struct {
	Addr         string
	Exchange     string //交换机名
	RoutingKey   string //路由键
	ExchangeType string //交换机类型
	QueueName    string //队列名
	AutoDelete   bool   //是否自动删除,在最后一个consumer断开连接后，删除
}
