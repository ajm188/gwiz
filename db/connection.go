package db

import (
	"database/sql"
)

type Connection interface {
	closable
	queryable
	preparable
	Begin() (Transaction, error)
	Ping() error
}

type Conn struct {
	*sql.DB
}

func (conn *Conn) Begin() (Transaction, error) {
	tx, err := conn.DB.Begin()
	return &Tx{tx}, err
}

func (conn *Conn) Prepare(query string) (Statement, error) {
	stmt, err := conn.DB.Prepare(query)
	return &Stmt{stmt}, err
}
