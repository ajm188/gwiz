package db

type Connection interface {
	closable
	queryable
	Begin() (Transaction, error)
	Prepare(string) (Statement, error)
	Ping() error
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

func Database(connStr *string) (Connection, error) {
	if connStr == nil {
		str := ConnectionString()
		connStr = &str
	}
	sqlDB, err := sql.Open("postgres", *connStr)
	return &DB{sqlDB}, err
}
