package main

import (
	"chat/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var (
	addr = ":9050"
	uniformDomain = 0.0002
	normDomain = 0.0002
	normMean = 0.00001
	oscillationPeriod = 10 * time.Minute
)

var (
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "rpc_durations_seconds",
			Help: "RPC latence distributions",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"service"})
)

func init()  {

	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())

}

func main()  {
	start := time.Now()

	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(oscillationPeriod)))
	}

	go func() {
		for {
			v := rand.Float64() * uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * normDomain) + normMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64() / 1e6
			rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(addr, nil))
}