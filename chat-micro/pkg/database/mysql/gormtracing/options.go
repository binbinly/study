package gormtracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type options struct {
	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool

	// tracer allows users to use customized and different tracer to makes tracing clearly.
	tracer trace.Tracer

	// Whether to log statement parameters or leave placeholders in the queries.
	logSqlParameters bool

	// errorTagHook allows users to customized error what kind of error tag should be tagged.
	errorTagHook errorTagHook
}

func defaultOption() *options {
	return &options{
		logResult:        false,
		tracer:           otel.Tracer("mysql-gorm"),
		logSqlParameters: true,
		errorTagHook:     defaultErrorTagHook,
	}
}

type ApplyOption func(o *options)

// WithLogResult enable opentracingPlugin to log the result of each executed sql.
func WithLogResult() ApplyOption {
	return func(o *options) {
		o.logResult = true
	}
}

// WithTracer allows to use customized tracer rather than the global one only.
func WithTracer(tracer trace.Tracer) ApplyOption {
	return func(o *options) {
		if tracer == nil {
			return
		}

		o.tracer = tracer
	}
}

func WithSqlParameters(logSqlParameters bool) ApplyOption {
	return func(o *options) {
		o.logSqlParameters = logSqlParameters
	}
}

func WithErrorTagHook(errorTagHook errorTagHook) ApplyOption {
	return func(o *options) {
		if errorTagHook == nil {
			return
		}

		o.errorTagHook = errorTagHook
	}
}
