package torm

import (
	"github.com/jackie8tao/torm/sqlx"
	"github.com/jackie8tao/torm/stmt"
)

type selectVisitor struct {
	builder    *SQLBuilder
	dialectOpt *SQLDialectOptions
}

func (s *selectVisitor) Visit(rawExpr sqlx.Stmt) (err error) {
	_, ok := rawExpr.(stmt.SelectStmt)
	if !ok {
		panic(sqlx.NewErr("invalid sql fragment"))
	}
	return
}
