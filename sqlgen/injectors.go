package sqlgen

import "git.code.oa.com/pluto/torm/expr"

// injector is a function that injects sql expression into builder.
type injector func(SQLBuilder) error

func injectExpr(builder SQLBuilder, expr expr.Expr) (err error) {
	return builder.WriteExpr(expr)
}

func injectExprs(builder SQLBuilder, exprs []expr.Expr) (err error) {
	for _, v := range exprs {
		err = builder.WriteExpr(v)
		if err != nil {
			return
		}
	}
	return
}
