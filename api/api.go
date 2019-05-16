package api

import "net/http"

func Run(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
