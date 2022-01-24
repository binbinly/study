package server

import (
	"chat-micro/app/logic/conf"
	"chat-micro/app/logic/routers"
	"chat-micro/pkg/transport/http"
)

// NewHTTPServer creates a HTTP er
func NewHTTPServer(c *conf.HTTPConfig) *http.Server {
	router := routers.NewRouter(true)

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router
	// NOTE: register svc to http server

	return srv
}
