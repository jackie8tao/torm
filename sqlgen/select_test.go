package sqlgen

import (
	"testing"

	"git.code.oa.com/pluto/torm/expr"
)

func TestSelectStmt_ToSQL(t *testing.T) {
	stmt := &SelectStmt{
		distinct: true,
		columns:  expr.NewColExpr("*"),
		from:     expr.NewFromExpr("tableA"),
		joins: expr.NewJoinList(
			expr.NewJoinExpr(expr.LeftJoin, "tableB", "tableA.a = tableB.a"),
			expr.NewJoinExpr(expr.RightJoin, "tableC", "tableC.a = tableB.a"),
		),
		wheres:   expr.NewWhereExpr(expr.AndOper, "a = b"),
		groupBys: expr.NewColList("a", "b", "c"),
		havings: expr.NewWhereList(
			expr.NewWhereExpr(expr.AndOper, "a=1"),
			expr.NewWhereExpr(expr.OrOper, "b=2"),
		),
		ordersBy: expr.NewOrderByList(
			expr.NewOrderByExpr(expr.DescOrder, "a"),
			expr.NewOrderByExpr(expr.AscOrder, "b"),
		),
		limit:  100,
		offset: 10,
	}
	sql, _, err := stmt.ToSQL()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
