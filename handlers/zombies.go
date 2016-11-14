package handlers

import (
	"fmt"
	"net/http"
)

import (
	"github.com/ajm188/gwiz/db"
)

func zombies() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/detail/", zombieDetail)
	mux.HandleFunc("/new/", zombieNew)
	mux.HandleFunc("/", WithConnectionFunc(zombieBase))

	return mux
}

func zombieBase(conn db.Connection, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		zombieIndex(conn, w, r)
	case "POST":
		zombieCreate(conn, w, r)
	default:
		fmt.Fprintf(w, "%s: %s\n", r.Method, r.RequestURI)
	}
}

func zombieDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from zombie detail!\n")
}

func zombieNew(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/zombies/new.html")
}

func zombieIndex(conn db.Connection, w http.ResponseWriter, r *http.Request) {
	var count int

	fmt.Fprintf(w, "Listing zombies\n")
	err := conn.QueryRow("SELECT COUNT(*) FROM zombies").Scan(&count)
	if err != nil {
		fmt.Fprintf(w, "Got %s while performing query\n", err)
		return
	}

	fmt.Fprintf(w, "%d\n", count)
}

func zombieCreate(conn db.Connection, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I see you are trying to create a zombie!\n")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Uh oh. Got %s when parsing form", err)
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "You asked to create a zombie with name: %s\n", name)

	stmt, err := conn.Prepare("INSERT INTO zombies (name) VALUES ($1) RETURNING id;")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		fmt.Fprintf(w, "That zombie was created with id %d\n", id)
	}
	fmt.Fprintf(w, "One day this will 302 you to that zombie's page\n")
}
