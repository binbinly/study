package conf

import "github.com/spf13/viper"

//DefaultConf 默认配置
func DefaultConf(v *viper.Viper) {
	v.SetDefault("logger", map[string]interface{}{
		"Development":       false,
		"DisableCaller":     false,
		"Encoding":          "json",
		"Level":             "INFO",
		"Name":              "chat",
		"Writers":           "file",
		"LoggerFile":        "./logs/chat.log",
		"LoggerWarnFile":    "./logs/chat.warn.log",
		"LoggerErrorFile":   "./logs/chat.err.log",
	})
	v.SetDefault("trace", map[string]interface{}{
		"ServiceName": "app",
		"TraceAgent":  "jaeger",
		"OpenDebug":   false,
		"Enable": false,
		"Zipkin": map[string]interface{}{
			"HTTPEndpoint": "http://127.0.0.1:9411/api/v2/spans",
			"SameSpan":     false,
			"ID128Bit":     true,
			"SampleRate":   1.0,
		},
		"Jaeger": map[string]interface{}{
			"SamplingServerURL":      "http://127.0.0.1:5778/sampling",
			"SamplingType":           "const",
			"SamplingParam":          1.0,
			"LocalAgentHostPort":     "127.0.0.1:6831",
			"Propagation":            "jaeger",
			"Gen128Bit":              true,
			"TraceContextHeaderName": "uber-trace-id",
			"CollectorEndpoint":      "",
			"CollectorUser":          "",
			"CollectorPassword":      "",
		},
		"Elastic": map[string]interface{}{
			"ServerURL":   "http://127.0.0.1:8200",
			"SecretToken": "",
		},
	})
}
