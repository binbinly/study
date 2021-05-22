package log

import (
	"context"
	"fmt"
)

// log is A global variable so that log functions can be directly accessed
var log Logger

// logger is A global variable with trace log
var logger Factory

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

// Logger config
type Config struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
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

func NewConfig() *Config {
	return &Config{
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
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
	zapLogger, err := newZapLogger(cfg)
	if err != nil {
		fmt.Errorf("Init newZapLogger err: %v", err)
	}
	l, err := newLogger(cfg)
	if err != nil {
		fmt.Errorf("Init newLogger err: %v", err)
	}

	// init logger with trace log
	logger = NewFactory(zapLogger, l)

	// normal log
	log = l

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

func GetLogger() Logger {
	return log
}

// Trace is a logger that can log msg and log span for trace
func Trace(ctx context.Context) Logger {
	return logger.For(ctx)
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
