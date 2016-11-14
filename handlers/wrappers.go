package handlers

import (
	"net/http"
)

import (
	"github.com/ajm188/gwiz/db"
)

type ConnectionHandlerFunc func(db.Connection, http.ResponseWriter, *http.Request)

func WithConnectionFunc(handlerFunc ConnectionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := db.NewConnection(nil)
		if err != nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		defer conn.Close()
		handlerFunc(conn, w, r)
	}
}
