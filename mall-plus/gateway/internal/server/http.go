package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"gateway/conf"
	"pkg/transport/http"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *conf.HTTPConfig, mux *runtime.ServeMux) *http.Server {
	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)
	srv.Handler = mux

	return srv
}
