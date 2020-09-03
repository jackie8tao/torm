package expr

import (
	"fmt"
)

// TableExpr from expression.
type TableExpr struct {
	table string
}

// ToSQL this function implements Expr interface.
func (f TableExpr) ToSQL() (sql string, args []interface{}, err error) {
	if f.table == "" {
		err = ErrEmptyTable
		return
	}
	sql = fmt.Sprintf(" `%s`", f.table)
	return
}
