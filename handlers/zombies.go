package handlers

import (
	"fmt"
	"net/http"
)

func zombies() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/detail/", zombieDetail)
	mux.HandleFunc("/", zombieBase)

	return mux
}

func zombieBase(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: %s", r.Method, r.RequestURI)
}

func zombieDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from zombie detail!")
}
