package sqlBuilder

import (
	"fmt"
	queryBuilder "user/pkg/query_builder"
)

type Stmt struct {
	// baseSQL is the base SQL statement
	baseSQL string

	// whereCounter is the counter for WHERE condition
	whereCounter uint

	// values is the slice of values
	values []interface{}

	// formattedSQL is the formatted SQL statement
	formattedSQL string
}

func NewStmt() queryBuilder.QueryBuilder {
	return &Stmt{}
}

func (s *Stmt) SetBaseSQL(baseSQL string) queryBuilder.QueryBuilder {
	s.baseSQL = baseSQL
	return s
}

func (s *Stmt) SetWhere(condition string, value interface{}) queryBuilder.QueryBuilder {
	condition = fmt.Sprintf("%s = $%d", condition, len(s.values)+1)

	if s.whereCounter == 0 {
		condition = fmt.Sprintf("WHERE %s", condition)
	} else {
		condition = fmt.Sprintf("AND %s", condition)
	}

	if s.formattedSQL == "" {
		s.formattedSQL = fmt.Sprintf("%s %s", s.baseSQL, condition)
	} else {
		s.formattedSQL = fmt.Sprintf("%s %s", s.formattedSQL, condition)
	}

	s.values = append(s.values, value)
	s.whereCounter++

	return s
}

func (s *Stmt) Generate() (string, []interface{}) {
	if s.formattedSQL == "" {
		return s.baseSQL, s.values
	}

	return s.formattedSQL, s.values
}
