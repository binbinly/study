package http

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"pkg/transport"
	"pkg/utils"
)

var (
	_ transport.Server   = (*Server)(nil)
	_ transport.Endpoint = (*Server)(nil)
)

// Server http server struct
type Server struct {
	*http.Server
	lis          net.Listener
	network      string
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	endpoint     *url.URL
}

// defaultServer return a default config server
func defaultServer() *Server {
	return &Server{
		network:      "tcp",
		address:      ":8080",
		readTimeout:  time.Second,
		writeTimeout: time.Second,
	}
}

// NewServer create a server
func NewServer(opts ...ServerOption) *Server {
	srv := defaultServer()
	// apply options
	for _, o := range opts {
		o(srv)
	}
	// NOTE: must set server
	srv.Server = &http.Server{
		Addr:         srv.address,
		ReadTimeout:  srv.readTimeout,
		WriteTimeout: srv.writeTimeout,
		Handler:      srv,
	}
	return srv
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.ServeHTTP(resp, req)
}

// Endpoint return a real address to registry endpoint.
// examples:
//   http://127.0.0.1:8080
func (s *Server) Endpoint() (*url.URL, error) {
	addr, err := utils.Extract(s.address, s.lis)
	if err != nil {
		return nil, err
	}
	s.endpoint = &url.URL{Scheme: "http", Host: addr}
	return s.endpoint, nil
}

// Start start a server
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis

	if _, err := s.Endpoint(); err != nil {
		return err
	}
	log.Printf("[HTTP] server is listening on: %s", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop server
func (s *Server) Stop(ctx context.Context) error {
	log.Print("[HTTP] server is stopping")
	return s.Shutdown(ctx)
}
