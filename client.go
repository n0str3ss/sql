package sql

import (
	"errors"
	"fmt"
)

type Client interface {
	Insert(query string, args ...interface{}) (r Result, err error)
}

type MySqlClient struct {
	db *DB
}

func (c *MySqlClient) Insert(query string, args ...interface{}) (r Result, err error) {
	defer func() {
		if recovery := recover(); recovery != nil {
			err = errors.New(
				fmt.Sprintf(
					"sql.MySqlClient: panic while inserting into table. reason: %v",
					recovery,
				),
			)
		}
	}()

	smtInsert, err := c.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer smtInsert.Close()

	insResult, err := smtInsert.Exec(args...)

	if err != nil {
		return nil, err
	}

	r = &MySqlResult{
		Result: insResult,
	}

	return r, nil
}

func NewMySqlClient(db *DB) *MySqlClient {
	return &MySqlClient{
		db: db,
	}
}
