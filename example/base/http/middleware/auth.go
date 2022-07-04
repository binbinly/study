package middleware

import "net/http"

func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("token")
		if name == "pi" {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
}