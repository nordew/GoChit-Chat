package queryBuilder

import (
	"user/internal/repository"
)

type QueryBuilder interface {
	BuildQuery(filter *repository.GetFilter) (string, []interface{}, error)
}
