package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

// DB holds database conneciton
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbCOnn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates db pool for postgers
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetConnMaxIdleTime(maxIdleDbCOnn)
	d.SetConnMaxLifetime(maxDbLifetime)
	d.SetMaxOpenConns(maxOpenDbConn)

	dbConn.SQL = d

	err = TestDB(d)

	if err != nil {
		return nil, err
	}
	return dbConn, nil

}

// TestDB tries to ping database
func TestDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
