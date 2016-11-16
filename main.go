package main

import (
	"github.com/ajm188/gwiz/db"
	"github.com/ajm188/gwiz/handlers"
)

func main() {
	if err := db.Connect(nil); err != nil {
		return
	}
	defer db.Disconnect()
	handlers.Serve(":8080")
}
