package limiter

import (
	"net/http"

	"golang.org/x/time/rate"
)

// rate of 2, 5 tokens
var limiter = rate.NewLimiter(2, 5)

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
