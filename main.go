package main

import (
	"github.com/ajm188/gwiz/handlers"
	"github.com/ajm188/gwiz/db"
)

func main() {
	db.Database(nil)
	handlers.Serve(":8080")
}
