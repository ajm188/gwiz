package handlers

import (
	"net/http"
)

import (
	"github.com/ajm188/gwiz/db"
)

type RequestFunc func(*Request)

func WithRequest(f RequestFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{
			w,
			r,
			nil,
		}
		f(&req)
	}
}

func WithConnection(f RequestFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{
			w,
			r,
			nil,
		}
		conn, err := db.NewConnection(nil)
		if err != nil {
			req.Error(500, err)
			return
		}
		defer conn.Close()
		req.Connection = conn
		f(&req)
	}
}

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
