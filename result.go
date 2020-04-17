package sql

import "database/sql"

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type MySqlResult struct {
	sql.Result
}

func (r *MySqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *MySqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
