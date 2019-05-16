package api

import "net/http"

func run(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
