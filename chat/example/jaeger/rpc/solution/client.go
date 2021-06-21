package main

import (
	"chat/example/jaeger/lib"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"net/url"
	"os"
)

func main()  {

	if len(os.Args) != 2 {
		panic("Error: expecting one argument")
	}

	tracer, closer := lib.Init("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)
}

func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	u := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, u)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(), opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	resp, err := lib.Do(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err)
	}

	helloStr := string(resp)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr))

	return helloStr
}

func printHello(ctx context.Context, helloStr string)  {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	u := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, u)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(), opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	if _, err = lib.Do(req); err != nil {
		ext.LogError(span, err)
		panic(err)
	}
}