package torm

// Sqlizer is the interface that wraps the ToSql method.
type Sqlizer interface {
	// ToSql returns a SQL representation of the Sqlizer, along with a slice of args
	// as passed to e.g. database/sql.Exec. It can also return an error.
	ToSQL() (string, []interface{}, error)
}

type SqlAssigner interface {
}
