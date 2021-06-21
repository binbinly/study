package main

import (
	"chat/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main()  {

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9050", nil))
}
