package handlers

import (
	"net/http"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/zombies/", http.StripPrefix("/zombies", zombies()))
	mux.HandleFunc("/", index)

	return mux
}

func Serve(addr string) error {
	return http.ListenAndServe(addr, server())
}
