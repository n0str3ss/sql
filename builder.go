package sql

type QueryBuilder interface {
	Query() Query
}
