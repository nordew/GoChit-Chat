package sqlBuilder

import (
	"errors"
	"fmt"
	"strings"
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

	for i, field := range b.fields {
		placeholder := fmt.Sprintf("$%d", i+1)

		if value := getField(filter, field); value != "" {
			if len(args) > 0 {
				sqlQuery += ` OR`
			}

			sqlQuery += fmt.Sprintf(" %s = %s", field, placeholder)
			args = append(args, value)
		}
	}

	if len(args) == 0 {
		return "", nil, errors.New("filter criteria not provided")
	}

	return sqlQuery, args, nil
}

func getField(filter *repository.GetFilter, field string) string {
	switch field {
	case "ID":
		return filter.ID
	case "Email":
		return filter.Email
	default:
		return ""
	}
}

func formatFields(fields []string) string {
	return strings.Join(fields, ", ")
}
