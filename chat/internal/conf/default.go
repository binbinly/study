package conf

import "github.com/spf13/viper"

func DefaultConf(v *viper.Viper)  {
	v.SetDefault("logger", map[string]interface{}{
		"Development": false,
		"DisableCaller": false,
		"DisableStacktrace": false,
		"Encoding": "json",
		"Level": "INFO",
		"Name": "chat",
		"Writers": "console",
		"LoggerFile": "./logs/chat.log",
		"LoggerWarnFile": "./logs/chat.wf.log",
		"LoggerErrorFile": "./logs/chat.err.log",
		"LogRollingPolicy": "daily",
		"LogRotateDate": 1,
		"LogRotateSize": 1,
		"LogBackupCount": 7,
	})
	v.SetDefault("jaeger", map[string]interface{}{
		"Host": "127.0.0.1:6831",
		"ServiceName": "REST_API",
		"LogSpans": false,
	})
	v.SetDefault("prometheus", map[string]interface{}{
		"Enable": true,
	})
	v.SetDefault("sentry", map[string]interface{}{
		"Enable": true,
	})
}