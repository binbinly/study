package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"

	"mall/pkg/net/ip"
)

const (
	// WriterConsole console输出
	WriterConsole = "console"
	// WriterFile 文件输出
	WriterFile = "file"
)

const (
	// RotateTimeDaily 按天切割
	RotateTimeDaily = "daily"
	// RotateTimeHourly 按小时切割
	RotateTimeHourly = "hourly"
)

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(cfg *Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// zapLogger logger struct
type zapLogger struct {
	sugarLogger *zap.SugaredLogger
}

// newZapLogger new zap logger
func newZapLogger(cfg *Config) (*zap.Logger, error) {
	return buildLogger(cfg), nil
}

// newLogger new logger
func newLogger(cfg *Config) (Logger, error) {
	return &zapLogger{sugarLogger: buildLogger(cfg).Sugar()}, nil
}

func buildLogger(cfg *Config) *zap.Logger {
	var encoderCfg zapcore.EncoderConfig
	if cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == WriterConsole {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	var cores []zapcore.Core
	var options []zap.Option
	// init option
	hostname, _ := os.Hostname()
	option := zap.Fields(
		zap.String("ip", ip.GetLocalIP()),
		zap.String("app_id", cfg.Name),
		zap.String("instance_id", hostname),
	)
	options = append(options, option)

	writers := strings.Split(cfg.Writers, ",")
	for _, w := range writers {
		switch w {
		case WriterConsole:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))
		case WriterFile:
			// info
			cores = append(cores, getInfoCore(encoder, cfg))

			// warning
			core, option := getWarnCore(encoder, cfg)
			cores = append(cores, core)
			if option != nil {
				options = append(options, option)
			}

			// error
			core, option = getErrorCore(encoder, cfg)
			cores = append(cores, core)
			if option != nil {
				options = append(options, option)
			}
		default:
			// console
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))
			// file
			cores = append(cores, getAllCore(encoder, cfg))
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	if !cfg.DisableCaller {
		caller := zap.AddCaller()
		options = append(options, caller)
	}

	// 跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(2)
	options = append(options, addCallerSkip)

	// 构造日志
	return zap.New(combinedCore, options...)
}

func getAllCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	allWriter := getLogWriterWithTime(cfg, cfg.LoggerFile)
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
}

func getInfoCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	infoWrite := getLogWriterWithTime(cfg, cfg.LoggerFile)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
}

func getWarnCore(encoder zapcore.Encoder, cfg *Config) (zapcore.Core, zap.Option) {
	warnWrite := getLogWriterWithTime(cfg, cfg.LoggerWarnFile)
	var stacktrace zap.Option
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if !cfg.DisableCaller {
			stacktrace = zap.AddStacktrace(zapcore.WarnLevel)
		}
		return lvl == zapcore.WarnLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(warnWrite), warnLevel), stacktrace
}

func getErrorCore(encoder zapcore.Encoder, cfg *Config) (zapcore.Core, zap.Option) {
	errorFilename := cfg.LoggerErrorFile
	errorWrite := getLogWriterWithTime(cfg, errorFilename)
	var stacktrace zap.Option
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if !cfg.DisableCaller {
			stacktrace = zap.AddStacktrace(zapcore.ErrorLevel)
		}
		return lvl >= zapcore.ErrorLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), errorLevel), stacktrace
}

// getLogWriterWithTime 打开获取文件句柄
func getLogWriterWithTime(cfg *Config, filename string) io.Writer {
	dirname := filepath.Dir(filename)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		log.Fatal(err)
	}
	// if we got here, then we need to create a file
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fh
}

// Debug logger
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Info logger
func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Warn logger
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Error logger
func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	logger := l.sugarLogger.With(f...)
	return &zapLogger{logger}
}
