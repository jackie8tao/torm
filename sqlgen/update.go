package sqlgen

import (
	"fmt"

	"git.code.oa.com/pluto/torm/expr"
)

// UpdateStmt update statement.
type UpdateStmt struct {
	table    expr.Expr
	values   expr.Expr
	wheres   expr.Expr
	ordersBy expr.Expr
	limit    int
}

// updateInjector
func (u *UpdateStmt) updateInjector(builder SQLBuilder) error {
	return builder.WriteString("UPDATE")
}

// tableInjector
func (u *UpdateStmt) tableInjector(builder SQLBuilder) error {
	if u.table == nil {
		return nil
	}
	return builder.WriteExpr(u.table)
}

// valuesInjector
func (u *UpdateStmt) valuesInjector(builder SQLBuilder) error {
	if u.wheres == nil {
		return nil
	}
	err := builder.WriteString(" SET")
	if err != nil {
		return err
	}
	return builder.WriteExpr(u.wheres)
}

// orderByInjector
func (u *UpdateStmt) orderByInjector(builder SQLBuilder) error {
	if u.ordersBy == nil {
		return nil
	}
	err := builder.WriteString(" ORDER BY")
	if err != nil {
		return err
	}
	return builder.WriteExpr(u.ordersBy)
}

// limitInjector
func (u *UpdateStmt) limitInjector(builder SQLBuilder) error {
	if u.limit <= 0 {
		return nil
	}
	return builder.WriteString(fmt.Sprintf(" LIMIT %d", u.limit))
}

// ToSQL returns the update sql statement and arguments.
// if it occurs errors, it will return the error.
func (u *UpdateStmt) ToSQL() (sql string, args []interface{}, err error) {
	return
}
