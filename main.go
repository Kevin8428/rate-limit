package main

import (
	"net/http"

	"github.com/Kevin8428/rate-limit/api"
	"github.com/Kevin8428/rate-limit/limiter"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Run)
	http.ListenAndServe(":8080", limiter.Limit(mux))
}
