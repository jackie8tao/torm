package expr

// Expr is the interface that all sql segments should implements.
type Expr interface {
	// ToSQL returns a SQL representation of the expression, along with a slice of args
	// as passed to e.g. database/sql.Exec. It can also return an error.
	ToSQL() (string, []interface{}, error)
}
