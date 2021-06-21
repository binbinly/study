package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func main()  {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		handler := new(cgi.Handler)

		handler.Path = ""

		handler.Dir = ""

		script := "testdata/cgi/" + request.URL.Path

		args := []string{"run", script}

		handler.Args = append(handler.Args, args...)

		handler.Env = append(handler.Env, "")

		handler.InheritEnv = []string{"HOME", "GOCACHE"}

		handler.Logger = nil

		handler.Root = ""

		handler.ServeHTTP(writer, request)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
