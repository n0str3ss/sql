package sql

type Query interface {
	/*
		returns the query as a string.
		It is up to the implementation itself to decide how this is generated
	*/
	String() string
	Args() []interface{}
}

type SqlQuery struct {
	query string
	args  []interface{}
}

func (q *SqlQuery) String() string {
	return q.query
}

func (q *SqlQuery) Args() []interface{} {
	return q.args
}

func NewSqlQuery(query string, args []interface{}) *SqlQuery {
	return &SqlQuery{
		query: query,
		args:  args,
	}
}
