package main

import (
	"github.com/ajm188/gwiz/handlers"
)

func main() {
	handlers.Initialize()
	handlers.Serve(":8080")
}
