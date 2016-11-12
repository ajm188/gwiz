package db

import (
	"fmt"
	"os"
	"database/sql"
)

import (
	_ "github.com/lib/pq"
)

const (
	CONNECTION_FORMAT = "%s://%s:%s@%s/%s"
)

type closable interface {
	Close() error
}

type queryable interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type Statement interface {
	closable
	Exec(...interface{}) (sql.Result, error)
	Query(...interface{}) (*sql.Rows, error)
	QueryRow(...interface{}) *sql.Row
}

type Transaction interface {
	queryable
	Commit() error
	Rollback() error
	Stmt(stmt Statement) Statement
}

type Connection interface {
	closable
	queryable
	Begin() (Transaction, error)
	Prepare(string) (Statement, error)
	Ping() error
}

func Database(connStr *string) (Connection, error) {
	if connStr == nil {
		str := ConnectionString()
		connStr = &str
	}
	fmt.Println(*connStr)
}

func getEnv(name, defaultValue string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return
}

func ConnectionString() string {
	dialect := "postgres"
	user := getEnv("GWIZ_USER", getEnv("USER", "gwiz"))
	pass := os.Getenv("GWIZ_PASS")
	host := getEnv("GWIZ_HOST", "localhost")
	dbName := getEnv("GWIZ_DB", "gwiz")

	return generateConnectionString(dialect, user, pass, host, dbName)
}

func generateConnectionString(dialect, user, pass, host, dbName string) string {
	return fmt.Sprintf(CONNECTION_FORMAT, dialect, user, pass, host, dbName)
}
