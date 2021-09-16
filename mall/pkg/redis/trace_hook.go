package redis

import (
	"context"

	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

//redis opentracing hook
var _ redis.Hook = (*TracingHook)(nil)

//TracingHook 链路追踪钩子
type TracingHook struct {
	tracer opentracing.Tracer
}

//NewTracingHook 实例化redis链路追踪钩子
func NewTracingHook() *TracingHook {
	return &TracingHook{tracer: opentracing.GlobalTracer()}
}

//BeforeProcess 执行命令前调用
func (t TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	span, c := opentracing.StartSpanFromContext(ctx, cmd.FullName())
	if span == nil {
		return ctx, nil
	}

	ext.DBType.Set(span, "redis")
	ext.SpanKind.Set(span, "client")
	ext.DBStatement.Set(span, rediscmd.CmdString(cmd))

	return c, nil
}

//AfterProcess 执行命令后调用
func (t TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	if err := cmd.Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.Finish()
	return nil
}

//BeforeProcessPipeline 执行管道命令前调用
func (t TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	summary, cmdsString := rediscmd.CmdsString(cmds)

	span, c := opentracing.StartSpanFromContext(ctx, "pipeline "+summary)
	if span == nil {
		return ctx, nil
	}

	ext.DBType.Set(span, "redis")
	ext.SpanKind.Set(span, "client")
	ext.DBStatement.Set(span, cmdsString)
	span.SetTag("db.redis.num_cmd", len(cmds))

	return c, nil
}

//AfterProcessPipeline 执行管道命令后调用
func (t TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}
	if err := cmds[0].Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.Finish()
	return nil
}

func recordError(ctx context.Context, span opentracing.Span, err error) {
	if err != redis.Nil {
		ext.LogError(span, err)
	}
}