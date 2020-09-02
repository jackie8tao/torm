package expr

import (
	"fmt"
)

// FromExpr from expression.
type FromExpr struct {
	table string
}

// ToSQL this function implements Expr interface.
func (f FromExpr) ToSQL() (sql string, args []interface{}, err error) {
	if f.table == "" {
		err = ErrEmptyTable
		return
	}
	sql = fmt.Sprintf("` %s`", f.table)
	return
}
