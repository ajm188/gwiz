package handlers

import (
	"fmt"
	"net/http"
)

import (
	"github.com/ajm188/gwiz/models"
)

func zombies() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: figure out how to wrap these more cleanly
	mux.HandleFunc("/detail/", WithRequest(zombieDetail))
	mux.HandleFunc("/new/", WithRequest(zombieNew))
	mux.HandleFunc("/", WithTransaction(zombieBase))

	return mux
}

func zombieBase(request *Request) {
	switch request.Method {
	case "GET":
		zombieIndex(request)
	case "POST":
		zombieCreate(request)
	default:
		request.Render(fmt.Sprintf("%s: %s\n", request.Method, request.RequestURI))
	}
}

func zombieDetail(request *Request) {
	request.Render("Hello from zombie detail!\n")
}

func zombieNew(request *Request) {
	request.RenderFile("./templates/zombies/new.html")
}

func zombieIndex(request *Request) {
	txn := request.Transaction
	rows, err := txn.Query("SELECT id, name FROM zombies")
	if err != nil {
		request.Error(500, err)
		return
	}
	defer rows.Close()
	zombies := make([]*models.Zombie, 0)
	for rows.Next() {
		zombie := new(models.Zombie)
		err := rows.Scan(&zombie.Id, &zombie.Name)
		if err != nil {
			request.Error(500, err)
			return
		}
		zombies = append(zombies, zombie)
	}

	request.RenderTemplate("./templates/zombies/index.html",
		struct {
			Zombies []*models.Zombie
			Count   int
		}{
			zombies,
			len(zombies),
		},
	)
}

func zombieCreate(request *Request) {
	request.Render("I see you are trying to create a zombie!\n")
	if err := request.ParseForm(); err != nil {
		request.Error(500, err)
		return
	}
	name := request.FormValue("name")
	request.Render(fmt.Sprintf("You asked to create a zombie with name: %s\n", name))

	txn := request.Transaction
	stmt, err := txn.Prepare("INSERT INTO zombies (name) VALUES ($1) RETURNING id;")
	if err != nil {
		request.Error(500, err)
		return
	}

	rows, err := stmt.Query(name)
	if err != nil {
		request.Error(500, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			request.Error(500, err)
			return
		}
		request.Render(fmt.Sprintf("That zombie was created with id %d\n", id))
	}
	request.Render("One day this will 302 you to that zombie's page\n")
}
