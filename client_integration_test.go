//+build integration

package sql_test

import (
	"github.com/n0str3ss/sql"
	"sync"
	"testing"
)

/**
To run this tests you need a local mysql (check docker-compose file and start mysql service)
after mysql container is running you have to run the sql queries located under docs dir. //ToDo: add script to bootstrap mysql automatically
You can run the queries with a mysql client, directly in the container, or by running docker commands.

//Todo: improve tests to generate test independent tables, tables should be created in the beginning and dropped at the end of each test (either use some vendor or build from scratch)
*/

func TestMySqlClient_Insert(t *testing.T) {
	query := "INSERT INTO unit_test (field_one, field_two)VALUES (?,?)"
	args := []interface{}{10, "ten"}

	db, err := sql.NewDB("mysql",
		"root:1@tcp(localhost:3307)/local_unit_test",
	)

	if err != nil {
		t.Errorf("could not create DB connection. %v", err.Error())
		return
	}

	defer db.Close()

	client := sql.NewMySqlClient(db)

	r, err := client.Insert(query, args...)

	if err != nil {
		t.Errorf("expected no error but got %v", err.Error())
		return
	}

	if r == nil {
		t.Error("expected result to not be nil, but it was")
		return
	}

	rowsAffected, err := r.RowsAffected()

	if err != nil {
		t.Errorf("expected no error but got %v", err.Error())
		return
	}

	expectedRowsAffected := int64(1)
	if rowsAffected != expectedRowsAffected {
		t.Errorf("expected %v rows to be affected, but got %v", expectedRowsAffected, rowsAffected)
		return
	}
}

func TestMySqlClient_InsertWithRoutine(t *testing.T) {
	db, err := sql.NewDB("mysql",
		"root:1@tcp(localhost:3307)/local_unit_test",
	)

	defer db.Close()

	if err != nil {
		t.Errorf("could not create DB connection. %v", err.Error())
		return
	}

	client := sql.NewMySqlClient(db)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		query := "INSERT INTO unit_test (field_one, field_two)VALUES (?,?)"
		args := []interface{}{i, "something here"}
		go func(i int, query string, args ...interface{}) {
			r, err := client.Insert(query, args...)

			if err != nil {
				t.Errorf("expected no error but got %v", err.Error())
				return
			}

			if r == nil {
				t.Error("expected result to not be nil, but it was")
				return
			}

			rowsAffected, err := r.RowsAffected()

			if err != nil {
				t.Errorf("expected no error but got %v", err.Error())
				return
			}

			expectedRowsAffected := int64(1)
			if rowsAffected != expectedRowsAffected {
				t.Errorf("expected %v rows to be affected, but got %v", expectedRowsAffected, rowsAffected)
				return
			}
			wg.Done()
		}(i, query, args...)
	}

	wg.Wait()
}
