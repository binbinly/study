package server

import (
	"github.com/gin-gonic/gin"

	"chat/app/chat/conf"
	"chat/pkg/app"
	"chat/pkg/net/http"
	"chat/pkg/registry"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *conf.Config, rs registry.Registry, engine *gin.Engine) *http.Server {
	srv := http.NewServer(&c.HTTP, app.WithHost(c.App.Host), app.WithID(c.App.ServerID),
		app.WithName(c.App.Name+"-http"), app.WithRegistry(rs))
	srv.Handler = engine
	srv.Start()
	return srv
}
