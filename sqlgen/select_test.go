package sqlgen

import (
	"testing"

	"git.code.oa.com/pluto/torm/expr"
)

func TestSelectStmt_ToSQL(t *testing.T) {
	stmt := &SelectStmt{
		distinct: true,
		columns:  expr.NewColExpr(""),
		from:     expr.NewFromExpr("test"),
		joins:    expr.JoinList{},
		wheres:   expr.WhereList{},
		groupBys: expr.ColList{},
		havings:  expr.WhereList{},
		ordersBy: expr.OrderByList{},
		limit:    100,
		offset:   10,
	}
	sql, _, err := stmt.ToSQL()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
