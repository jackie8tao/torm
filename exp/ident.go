package exp

import "github.com/jackie8tao/torm/sqlx"

// IdentifierExp interface for identifierExp expression.
type IdentifierExp interface {
	Aliaseable
	Comparable
	Inable
	Isable
	Likeable
	Orderable
	Distinctable
	Castable
}

// sql expression for identifierExp
type identifierExp struct {
	schema string
	table  string
	col    interface{}
}

// NewIdentExp
func NewIdentExp(schema, table string, col interface{}) sqlx.Exp {
	return identifierExp{
		schema: schema,
		table:  table,
		col:    col,
	}
}

// Type get the type value of identifierExp
func (i identifierExp) Type() sqlx.ExpType {
	return sqlx.IdentifierExp
}

// Visit
func (i identifierExp) Visit(sb sqlx.SQLBuilder) (err error) {
	return
}
