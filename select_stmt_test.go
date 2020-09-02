package torm

import (
	"testing"
)

func TestSelectBuilder_ToSQL(t *testing.T) {
	sb := &SelectBuilder{
		distinct: false,
		columns: ColList{
			cols: []ColExpr{
				{table: "", column: "a"},
			},
		},
		fromParts: []Sqlizer{
			FromExpr{table: "test"},
		},
		joins:       nil,
		whereParts:  nil,
		groupBys:    nil,
		havingParts: nil,
		orderBys:    nil,
		limit:       0,
		offset:      0,
	}
	sql, _, err := sb.ToSQL()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
