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
	CONNECTION_FORMAT = "%s://%s:%s@%s/%s?sslmode=%s"
)

var (
	database *Conn
)

func Begin() (Transaction, error) {
	return database.Begin()
}

func Connect(connStr *string) error {
	checkConn := func(c *Conn) error {
		return c.Ping()
	}

	if database != nil {
		return checkConn(database)
	}

	if connStr == nil {
		str := ConnectionString()
		connStr = &str
	}
	sqlDB, err := sql.Open("postgres", *connStr)
	if err != nil {
		return err
	}
	database = &Conn{sqlDB}
	return checkConn(database)
}

func Disconnect() error {
	return database.Close()
}

func ConnectionString() string {
	dialect := "postgres"
	user := getEnv("GWIZ_USER", getEnv("USER", "gwiz"))
	pass := os.Getenv("GWIZ_PASS")
	host := getEnv("GWIZ_HOST", "localhost")
	dbName := getEnv("GWIZ_DB", "gwiz")
	sslMode := getEnv("GWIZ_USE_SSL", "disable")

	return generateConnectionString(dialect, user, pass, host, dbName, sslMode)
}

type closable interface {
	Close() error
}

type queryable interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type preparable interface {
	Prepare(string) (Statement, error)
}

func getEnv(name, defaultValue string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return
}

func generateConnectionString(dialect, user, pass, host, dbName, sslMode string) string {
	return fmt.Sprintf(CONNECTION_FORMAT, dialect, user, pass, host, dbName, sslMode)
}
