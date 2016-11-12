package db

import (
	"database/sql"
)

type Statement interface {
	closable
	Exec(...interface{}) (sql.Result, error)
	Query(...interface{}) (*sql.Rows, error)
	QueryRow(...interface{}) *sql.Row
}

type Stmt struct {
	*sql.Stmt
}
