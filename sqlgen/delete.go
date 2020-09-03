package sqlgen

import "git.code.oa.com/pluto/torm/expr"

// DeleteStmt delete sql statement.
type DeleteStmt struct {
	table  expr.Expr
	wheres expr.Expr
	limit  int
}
