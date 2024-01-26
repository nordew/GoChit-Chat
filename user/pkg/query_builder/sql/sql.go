package sqlBuilder

import (
	"errors"
	"fmt"
	"user/internal/repository"
)

type sqlBuilder struct {
	table  string
	fields []string
}

func NewSqlBuilder(table string, fields []string) *sqlBuilder {
	return &sqlBuilder{
		table:  table,
		fields: fields,
	}
}

func (b *sqlBuilder) BuildDynamicQuery(filter *repository.GetFilter) (string, []interface{}, error) {
	if filter == nil {
		return "", nil, errors.New("filter is nil")
	}

	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE", formatFields(b.fields), b.table)
	var args []interface{}

	if filter.ID != "" {
		sqlQuery += ` id = $1`
		args = append(args, filter.ID)
	}

	if filter.Email != "" {
		if len(args) > 0 {
			sqlQuery += ` OR`
		}
		sqlQuery += ` email = $2`
		args = append(args, filter.Email)
	}

	if len(args) == 0 {
		return "", nil, errors.New("filter criteria not provided")
	}

	return sqlQuery, args, nil
}

func formatFields(fields []string) string {
	return fmt.Sprintf("%s", fields)
}
