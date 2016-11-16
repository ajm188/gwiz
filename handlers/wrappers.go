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

func WithTransaction(f RequestFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{
			w,
			r,
			nil,
		}
		txn, err := db.Begin()
		if err != nil {
			req.Error(500, err)
			return
		}
		req.Transaction = txn
		f(&req)

		if err = txn.Commit(); err != nil {
			req.Error(500, err)
		}
	}
}
