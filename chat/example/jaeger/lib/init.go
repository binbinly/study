package lib

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func Init(service string) (opentracing.Tracer, io.Closer) {

	cfg := &config.Configuration{
		ServiceName:         service,
		Sampler: &config.SamplerConfig{
			SamplingServerURL:"http://192.168.8.76:5778/sampling",
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: "192.168.8.76:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}