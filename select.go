package sql

type ConditionQuery struct {
}

type Where struct {
	field    string
	operator string
	value    string
}

func (c *ConditionQuery) Where(condition *Where) *ConditionQuery {
	return c
}

func (c *ConditionQuery) And(condition *Where) *ConditionQuery {
	return c
}

type SelectQuery struct {
	*ConditionQuery
	table         string
	sort          string
	sortDirection string
}

func Select() *SelectQuery {
	return &SelectQuery{}
}

func (s *SelectQuery) From(table string) *SelectQuery {
	s.table = table

	return s
}

func (s *SelectQuery) Sort(field, direction string) *SelectQuery {
	s.sort = field
	s.sortDirection = direction

	return s
}

func (s *SelectQuery) Query() Query {
	return NewSqlQuery("", nil)
}
