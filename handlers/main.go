package handlers

import (
	"net/http"
)

func Initialize() {
	http.HandleFunc("/", index)
}

func Serve(addr string) error {
	return http.ListenAndServe(addr, nil)
}
