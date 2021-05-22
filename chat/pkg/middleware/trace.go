package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"chat/pkg/log"
)

const (
	// DefaultServiceName service name
	DefaultServiceName = "snake"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()

		var sp opentracing.Span
		// for http
		spanCtx, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Warn("err", err)
		}
		sp = tracer.StartSpan(
			"HTTP "+c.Request.Method+" "+c.Request.URL.Path,
			ext.RPCServerOption(spanCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		)

		// record HTTP method
		ext.HTTPMethod.Set(sp, c.Request.Method)
		// record HTTP url
		ext.HTTPUrl.Set(sp, c.Request.URL.String())

		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), sp))

		c.Next()

		// record HTTP status code
		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))

		sp.Finish()
	}
}
