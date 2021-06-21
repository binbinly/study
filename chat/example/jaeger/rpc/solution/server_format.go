package main

import (
	"chat/example/jaeger/lib"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func main()  {

	tracer, closer := lib.Init("formatter")
	defer closer.Close()

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("hello, %s!", helloTo)
		span.LogFields(
			log.String("event", "string-format"),
			log.String("value", helloStr))
		w.Write([]byte(helloStr))
	})

	http.ListenAndServe(":8081", nil)

}