package main

import (
	"github.com/ajm188/gwiz/db"
	"github.com/ajm188/gwiz/handlers"
)

func main() {
	db.NewConnection(nil)
	handlers.Serve(":8080")
}
