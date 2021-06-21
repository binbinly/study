package main

import (
	"chat/example/jaeger/lib"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"os"
)

func main()  {

	if len(os.Args) != 2 {
		panic("Error: Expecting one argument")
	}

	tracer, closer := lib.Init("Hello-World")
	defer closer.Close()

	helloTo := os.Args[1]
	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
		)
	println(helloStr)

	span.LogKV("event", "println")
	span.Finish()
}
