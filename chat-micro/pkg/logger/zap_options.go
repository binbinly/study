package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) Option {
	return SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) Option {
	return SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) Option {
	return SetOption(namespaceKey{}, namespace)
}

type optionsKey struct{}

func WithOptions(opts ...zap.Option) Option {
	return SetOption(optionsKey{}, opts)
}
