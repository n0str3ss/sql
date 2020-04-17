package sql_test

import (
	"github.com/n0str3ss/sql"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name           string
		table          string
		fields         []string
		values         []interface{}
		expectedString string
		expectedArgs   []interface{}
	}{
		{
			name:  "test 1: expects to successfully create a query with one value",
			table: "table_name",
			fields: []string{
				"field_one",
			},
			values:         []interface{}{"valueOne"},
			expectedString: "INSERT INTO table_name (field_one) VALUES (?)",
			expectedArgs:   []interface{}{"valueOne"},
		},
		{
			name:  "test 2: expects to successfully create a query with two values",
			table: "table_name",
			fields: []string{
				"field_one",
				"field_two",
			},
			values:         []interface{}{"valueOne", 1},
			expectedString: "INSERT INTO table_name (field_one, field_two) VALUES (?, ?)",
			expectedArgs:   []interface{}{"valueOne", 1},
		},
		{
			name:  "test 3: expects to successfully create a query with ten values",
			table: "table_name",
			fields: []string{
				"field_one",
				"field_two",
				"field_three",
				"field_four",
				"field_five",
				"field_six",
				"field_seven",
				"field_eight",
				"field_nine",
				"field_ten",
			},
			values: []interface{}{
				"valueOne",
				1,
				true,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
			},
			expectedString: "INSERT INTO table_name (field_one, field_two, field_three, field_four, field_five, field_six, field_seven, field_eight, field_nine, field_ten) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			expectedArgs: []interface{}{
				"valueOne",
				1,
				true,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
			},
		},
	}

	for _, test := range tests {
		ins := sql.Insert().Into(test.table).Fields(test.fields).Values(test.values)

		actualString := ins.Query().String()
		actualArgs := ins.Query().Args()

		if actualString != test.expectedString {
			t.Errorf("%v: wrong query string received. want\n %v\n got\n %v\n", test.name, test.expectedString, actualString)
		}

		if reflect.DeepEqual(test.expectedArgs, actualArgs) == false {
			t.Errorf("%v: wrong query args received. want\n %v\n got\n %v\n", test.name, test.expectedArgs, actualArgs)
		}
	}
}
