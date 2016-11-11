package main

import (
	"github.com/ajm188/gwiz/views"
)

func main() {
	views.Initialize()
	views.Serve(":8080")
}
