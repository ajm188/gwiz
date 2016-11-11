package handlers

import (
	"fmt"
	"net/http"
)

func zombies() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/detail/", zombieDetail)
	mux.HandleFunc("/new/", zombieNew)
	mux.HandleFunc("/", zombieBase)

	return mux
}

func zombieBase(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "I see you are trying to create a zombie!\n")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Uh oh. Got %s when parsing form", err)
			return
		}
		name := r.FormValue("name")
		fmt.Fprintf(w, "You asked to create a zombie with name: %s", name)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", r.Method, r.RequestURI)
}

func zombieDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from zombie detail!\n")
}

func zombieNew(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/zombies/new.html")
}
