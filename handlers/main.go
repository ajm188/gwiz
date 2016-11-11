package handlers

import (
	"net/http"
)

func Initialize() {
	http.Handle("/zombies/", zombies())
	http.HandleFunc("/", index)
}

func Serve(addr string) error {
	return http.ListenAndServe(addr, nil)
}
