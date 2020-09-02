package expr

import (
	"fmt"
)

const (
	LeftJoin JoinMethod = iota + 1
	RightJoin
	InnerJoin
)

// JoinMethod join type, eg: LEFT JOIN, RIGHT JOIN and so on.
type JoinMethod int

// JoinExpr join expression.
type JoinExpr struct {
	method JoinMethod
	table  string
	cond   string
}

// ToSQL this function implements Expr interface.
func (j JoinExpr) ToSQL() (sql string, args []interface{}, err error) {
	var method string
	switch j.method {
	case LeftJoin:
		method = "LEFT JOIN"
	case RightJoin:
		method = "RIGHT JOIN"
	case InnerJoin:
		method = "INNER JOIN"
	default:
		err = ErrIllegalJoinMethod
		return
	}

	sql = fmt.Sprintf(" %s %s ON %s", method, j.table, j.cond)

	return
}

// JoinList join expression list.
type JoinList struct {
	joins []JoinExpr
}

// ToSQL this function implements Expr interface.
func (j JoinList) ToSQL() (sql string, args []interface{}, err error) {
	var seg string
	for _, v := range j.joins {
		seg, _, err = v.ToSQL()
		if err != nil {
			return
		}
		sql += seg
	}

	return
}
