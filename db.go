package sql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func (db *DB) Close() {
	db.DB.Close()
}

func NewDB(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("sql could not open connection to db. reason: %v", err.Error()))
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("sql could not ping db. reason: %v", err.Error()))
	}
	return &DB{db}, nil
}
