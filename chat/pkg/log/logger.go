package log

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log is A global variable so that log functions can be directly accessed
var log Logger
var zl *zap.Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//Config logger
type Config struct {
	Development       bool
	DisableCaller     bool
	Encoding          string
	Level             string
	Name              string
	Writers           string
	LoggerFile        string
	LoggerWarnFile    string
	LoggerErrorFile   string
	LogFormatText     bool
	LogRollingPolicy  string
	LogRotateDate     int
	LogRotateSize     int
	LogBackupCount    uint
}

//NewConfig 实例化默认配置
func NewConfig() *Config {
	return &Config{
		Development:       false,
		DisableCaller:     false,
		Encoding:          "json",
		Level:             "info",
		Name:              "app",
		Writers:           "console",
		LoggerFile:        "./logs/app.log",
		LoggerWarnFile:    "./logs/warn.log",
		LoggerErrorFile:   "./logs/error.log",
		LogFormatText:     false,
		LogRollingPolicy:  "daily",
		LogRotateDate:     1,
		LogRotateSize:     1,
		LogBackupCount:    7,
	}
}

// InitLog init log
func InitLog(cfg *Config) Logger {
	var err error
	// new zap logger
	zl, err = newZapLogger(cfg)
	if err != nil {
		fmt.Errorf("init newZapLogger err: %v", err)
	}
	_ = zl

	// new sugar logger
	log, err = newLogger(cfg)
	if err != nil {
		fmt.Errorf("init newLogger err: %v", err)
	}

	return log
}

// Logger is our contract for the logger
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

//GetLogger 获取log对象
func GetLogger() Logger {
	return log
}

//WithContext Trace is a logger that can log msg and log span for trace
func WithContext(ctx context.Context) Logger {
	//return logger.For(ctx)
	if span := opentracing.SpanFromContext(ctx); span != nil {
		logger := spanLogger{span: span, logger: zl}

		if jaegerCtx, ok := span.Context().(jaeger.SpanContext); ok {
			logger.spanFields = []zapcore.Field{
				zap.String("trace_id", jaegerCtx.TraceID().String()),
				zap.String("span_id", jaegerCtx.SpanID().String()),
			}
		}

		return logger
	}
	return log
}

// Debug logger
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logger
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logger
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logger
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal logger
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf logger
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf logger
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields logger
// output more field, eg:
// 		contextLogger := log.WithFields(log.Fields{"key1": "value1"})
// 		contextLogger.Info("print multi field")
// or more sample to use:
// 	    log.WithFields(log.Fields{"key1": "value1"}).Info("this is a test log")
// 	    log.WithFields(log.Fields{"key1": "value1"}).Infof("this is a test log, user_id: %d", userID)
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
