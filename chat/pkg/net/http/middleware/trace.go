package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

//Trace 链路追踪中间件
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()

		var newCtx context.Context
		var sp opentracing.Span
		// for http
		spanCtx, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			// root
			sp, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer,
				c.Request.URL.Path,
			)
		} else {
			sp, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}

		// record HTTP method
		ext.HTTPMethod.Set(sp, c.Request.Method)
		// record HTTP url
		ext.HTTPUrl.Set(sp, c.Request.URL.String())
		ext.PeerHostIPv4.SetString(sp, c.ClientIP())
		ext.PeerHostname.Set(sp, c.Request.Host)

		// add trace id and span id
		// get trace id and span id by using log
		var traceID string
		var spanID string
		var spanContext = sp.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)

		c.Request = c.Request.WithContext(newCtx)

		c.Next()

		// record HTTP status code
		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))

		sp.Finish()
	}
}
