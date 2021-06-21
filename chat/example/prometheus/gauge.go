package main

import (
	"chat/example/prometheus/metrics"
	"chat/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors",
		},[]string{"device", "service"})
)

func init()  {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main()  {
	go func() {
		for {
			val := rand.Float64() * 100
			cpuTemp.Set(val)
			hdFailures.With(prometheus.Labels{
				"device": "/dev/sda",
				"service": "hello.world",
			}).Inc()
			time.Sleep(time.Second)
		}
	}()
	gauge()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9050", nil))

}

func gauge()  {
	vec := metrics.NewGaugeVec(&metrics.GaugeVecOpts{
		Namespace: "rpc_client2",
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "rpc server requests duration(ms).",
		Labels:    []string{"path"},
	})

	go func() {
		for {
			vec.Inc("/test_inc")
			vec.Add(30, "/add")
			vec.Set(666, "/set")
			time.Sleep(time.Second * 2)
		}
	}()
}