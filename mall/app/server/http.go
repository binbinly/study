package server

import (
	"github.com/gin-gonic/gin"

	"mall/app/conf"
	"mall/pkg/net/http"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *conf.Config, engine *gin.Engine) *http.Server {
	srv := http.NewServer(&c.HTTP)
	srv.Handler = engine
	srv.Start()
	return srv
}
