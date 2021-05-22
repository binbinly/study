package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"chat/app/logic/conf"
	"chat/pkg/registry"
)

// StartServer start http server
func StartServer(rs registry.Registry) *http.Server {
	router := NewRouter(conf.Conf.App.Debug)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Conf.Http.Port),
		Handler:      router,
		ReadTimeout:  conf.Conf.Http.ReadTimeout,
		WriteTimeout: conf.Conf.Http.WriteTimeout,
	}

	log.Printf("Listening and serving HTTP on %d\n", conf.Conf.Http.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe, err: %s", err.Error())
		}
	}()

	// 服务注册
	err := rs.Register(context.Background(), &registry.Service{
		Id:   "h-" + conf.Conf.App.ServerId,
		Name: conf.Conf.App.Name + "-http",
		IP:   conf.Conf.App.Host,
		Port: conf.Conf.Http.Port,
		Check: registry.Check{
			HTTP: fmt.Sprintf("http://%v:%v/health", conf.Conf.App.Host, conf.Conf.Http.Port),
		},
	})
	if err != nil {
		log.Fatalf("[RegisterServer] failed to http register %s server: %v", conf.Conf.App.Name, err)
	}
	return srv
}
