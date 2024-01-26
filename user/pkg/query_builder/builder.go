package queryBuilder

type QueryBuilder interface {
	// SetBaseSQL sets the base SQL statement
	SetBaseSQL(baseSQL string) QueryBuilder

	// SetWhere sets the WHERE condition
	SetWhere(condition string, value interface{}) QueryBuilder

	// Generate generates the SQL statement
	Generate() (string, []interface{})
}
