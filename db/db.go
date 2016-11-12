package db

import (
	"database/sql"
	"fmt"
	"os"
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

type DB struct {
	*sql.DB
}

func (d *DB) Begin() (Transaction, error) {
	tx, err := d.DB.Begin()
	return &Tx{tx}, err
}

func (d *DB) Prepare(query string) (Statement, error) {
	stmt, err := d.DB.Prepare(query)
	return &Stmt{stmt}, err
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
	sqlDB, err := sql.Open("postgres", *connStr)
	return &DB{sqlDB}, err
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
