package conf

//DFSConfig 图片资源配置
type DFSConfig struct {
	Endpoint string `json:"endpoint"`
	Bucket   string `json:"bucket"`
}

//AMQPConfig amqp配置
type AMQPConfig struct {
	Addr string `json:"addr"`
}