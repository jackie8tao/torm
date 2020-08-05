package torm

import (
	"log"
	"testing"
)

func TestSelectStmt_ToSQL(t *testing.T) {
	where := &WhereNode{
		BaseNode: BaseNode{
			Type: NTWhere,
		},
		cond: "a > 100",
	}

	from := &FromNode{
		BaseNode: BaseNode{
			Type: NTFrom,
			Next: where,
		},
		table: "test",
	}

	st := &SelectNode{
		BaseNode: BaseNode{
			Type: NTSelect,
			Next: from,
		},
		columns: []string{"a", "b", "c", "d"},
	}

	s := &SelectStmt{
		Root: st,
	}

	sql, err := s.ToSQL()
	if err != nil {
		t.Error(err)
	}
	log.Println(sql)
}
