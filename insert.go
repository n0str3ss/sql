package sql

import (
	"fmt"
	"strings"
)

type InsertQuery struct {
	table  string
	fields []string
	values []interface{}
}

func Insert() *InsertQuery {
	return &InsertQuery{}
}

func (i *InsertQuery) Into(table string) *InsertQuery {
	i.table = table
	return i
}

func (i *InsertQuery) Fields(f []string) *InsertQuery {
	i.fields = f
	return i
}

func (i *InsertQuery) Values(v []interface{}) *InsertQuery {
	i.values = v
	return i
}

func (i *InsertQuery) Query() Query {
	return NewSqlQuery(i.generateQuerySting(), i.values)
}

func (i *InsertQuery) generateQuerySting() string {
	fieldsString := strings.Join(i.fields, ", ")

	placeholders := ""
	for range i.values {
		if placeholders != "" {
			placeholders += ", "
		}
		placeholders += "?"
	}

	return fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", i.table, fieldsString, placeholders)
}
