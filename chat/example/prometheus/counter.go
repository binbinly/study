package main

import (
	"chat/example/prometheus/metrics"
	"chat/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func recordMetrics()  {

	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name:"chat_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main()  {
	counter()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9050", nil))

}

func counter()  {
	vec := metrics.NewCounterVec(&metrics.CounterVecOpts{
		Namespace: "http_client",
		Subsystem: "call",
		Name:      "code_total",
		Help:      "http client requests error count.",
		Labels:    []string{"path", "code"},
	})
	go func() {
		for {
			vec.Inc("/test", "500")
			time.Sleep(2 * time.Second)
		}
	}()
}