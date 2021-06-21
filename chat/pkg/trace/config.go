package trace

import (
	"fmt"

	"chat/pkg/trace/elastic"
	"chat/pkg/trace/jaeger"
	"chat/pkg/trace/zipkin"
)

var (
	supportedTraceAgent = map[string]bool{
		zipkin.Name:  true,
		jaeger.Name:  true,
		elastic.Name: true,
	}
)

//Config 链路追踪配置
type Config struct {
	Enable      bool // 是否开启分布式追踪
	ServiceName string // The name of this service
	TraceAgent  string // The type of trace agent: zipkin, jaeger or elastic

	Zipkin  zipkin.Config  // Settings for zipkin, only useful when TraceAgent is zipkin
	Jaeger  jaeger.Config  // Settings for jaeger, only useful when TraceAgent is jaeger
	Elastic elastic.Config // Settings for elastic, only useful when TraceAgent is elastic
}

// Check 验证配置是否设置
func (cfg *Config) Check() error {

	if len(cfg.TraceAgent) == 0 {
		return fmt.Errorf("ModTrace.TraceAgent not set")
	}

	if _, ok := supportedTraceAgent[cfg.TraceAgent]; !ok {
		return fmt.Errorf("Trace.TraceAgent %s is not supported", cfg.TraceAgent)
	}

	return nil
}

// GetTraceConfig 获取跟踪配置
func (cfg *Config) GetTraceConfig() Agent {
	switch cfg.TraceAgent {
	case jaeger.Name:
		return &cfg.Jaeger
	case zipkin.Name:
		return &cfg.Zipkin
	case elastic.Name:
		return &cfg.Elastic
	default:
		return &cfg.Jaeger
	}
}
