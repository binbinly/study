package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

var (
	ipLimitMaps = make(map[string]*rate.Limiter)
	mu sync.Mutex
	rateLimit = 1
	rateMax = 5
)

func GetIPLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, ok := ipLimitMaps[ip]; ok {
		return limiter
	}
	limiter := rate.NewLimiter(rate.Limit(rateLimit), rateMax)
	ipLimitMaps[ip] = limiter

	return limiter
}

func IPRateLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := GetIPLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		handler.ServeHTTP(w, r)
	})
}