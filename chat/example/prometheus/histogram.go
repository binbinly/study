package main

import (
	"chat/example/prometheus/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func main() {
	histogram()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9050", nil))
}

func histogram()  {
	vec := metrics.NewHistogramVec(&metrics.HistogramVecOpts{
		Name:    "his_counts",
		Help:    "rpc server requests duration(ms).",
		Buckets: []float64{1, 2, 3},
		Labels:  []string{"method"},
	})
	go func() {
		for {
			vec.Observe(4, "/v1/users")
			time.Sleep(time.Second * 2)
		}
	}()

}