package main

import (
	"net/http"

	"github.com/Kevin8428/rate-limit/api"
)

func main() {
	http.HandleFunc("/run", api.Run)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
