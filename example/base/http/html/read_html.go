package main

import "net/http"

func main()  {
	dir := http.Dir("base/http/html/root")
	staticHandler := http.FileServer(dir)
	http.Handle("/", http.StripPrefix("/", staticHandler))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
