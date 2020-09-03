package sqlgen

import (
	"strings"

	"git.code.oa.com/pluto/torm/expr"
)

// injector is a function that injects sql expression into builder.
type injector func(SQLBuilder) error

// SQLBuilder
type SQLBuilder interface {
	expr.Expr
	WriteString(string) error
	WriteExpr(expr.Expr) error
}

// bufSQLBuilder
type bufSQLBuilder struct {
	buf  *strings.Builder
	args []interface{}
}

// WriteString appends string to the builder.
func (b *bufSQLBuilder) WriteString(str string) error {
	_, err := b.buf.WriteString(str)
	return err
}

// WriteExpr appends sql expression to the builder.
func (b *bufSQLBuilder) WriteExpr(expr expr.Expr) error {
	sql, args, err := expr.ToSQL()
	if err != nil {
		return err
	}
	b.args = append(b.args, args...)
	return b.WriteString(sql)
}

// ToSQL returns formatted sql string.
func (b *bufSQLBuilder) ToSQL() (sql string, args []interface{}, err error) {
	sql = b.buf.String()
	args = b.args
	return
}
