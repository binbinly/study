package redis

import (
	"context"

	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

//redis opentracing hook
var _ redis.Hook = (*TracingHook)(nil)

//TracingHook 链路追踪钩子
type TracingHook struct {
	tracer trace.Tracer
}

//NewTracingHook 实例化redis链路追踪钩子
func NewTracingHook() *TracingHook {
	return &TracingHook{tracer: otel.Tracer("redis")}
}

//BeforeProcess 执行命令前调用
func (t TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	var span trace.Span
	ctx, span = t.tracer.Start(ctx, cmd.FullName(),
		trace.WithSpanKind(trace.SpanKindClient))
	if span == nil {
		return ctx, nil
	}
	span.SetAttributes(
		semconv.DBSystemKey.String("redis"),
		semconv.DBStatementKey.String(rediscmd.CmdString(cmd)))

	return ctx, nil
}

//AfterProcess 执行命令后调用
func (t TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	if err := cmd.Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.End()
	return nil
}

//BeforeProcessPipeline 执行管道命令前调用
func (t TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	summary, cmdsString := rediscmd.CmdsString(cmds)
	var span trace.Span
	ctx, span = t.tracer.Start(ctx, "pipeline "+summary,
		trace.WithSpanKind(trace.SpanKindClient))
	if span == nil {
		return ctx, nil
	}
	span.SetAttributes(
		semconv.DBSystemKey.String("redis"),
		semconv.DBStatementKey.String(cmdsString),
		attribute.Key("db.redis.num_cmd").Int(len(cmds)))

	return ctx, nil
}

//AfterProcessPipeline 执行管道命令后调用
func (t TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	if err := cmds[0].Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.End()
	return nil
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != redis.Nil {
		span.RecordError(err)
	}
}
