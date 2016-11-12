package db

import (
	"database/sql"
)

type Transaction interface {
	queryable
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
