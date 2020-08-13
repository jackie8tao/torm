package stmt

import "github.com/jackie8tao/torm/sqlx"

type (
	BoolStmt interface {
		sqlx.Stmt
		Op()
		LHS()
		RHS()
	}
	boolStmt struct {
	}
)
