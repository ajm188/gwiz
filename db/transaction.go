package db

import (
	"database/sql"
)

type Transaction interface {
	queryable
	preparable
	Commit() error
	Rollback() error
	Stmt(stmt *Stmt) Statement
}

type Tx struct {
	*sql.Tx
}

func (tx *Tx) Stmt(stmt *Stmt) Statement {
	return &Stmt{tx.Tx.Stmt(stmt.Stmt)}
}

func (tx *Tx) Prepare(query string) (Statement, error) {
	stmt, err := tx.Tx.Prepare(query)
	return &Stmt{stmt}, err
}
