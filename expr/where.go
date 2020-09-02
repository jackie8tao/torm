package expr

import (
	"fmt"
)

// operators
const (
	AndOper OperVal = iota + 1
	OrOper
)

// OperVal operator value.
type OperVal int

// WhereExpr where expression.
type WhereExpr struct {
	oper OperVal
	cond string
}

// ToSQL this function implements Expr interface.
func (w WhereExpr) ToSQL() (sql string, args []interface{}, err error) {
	var oper string
	switch w.oper {
	case AndOper:
		oper = " AND"
	case OrOper:
		oper = " OR"
	default:
		err = ErrIllegalWhereOperator
		return
	}

	sql = fmt.Sprintf(" %s (%s)", oper, w.cond)
	return
}

// WhereList where list expression.
type WhereList struct {
	wheres []WhereExpr
}

// ToSQL this function implements Expr interface.
func (w WhereList) ToSQL() (sql string, args []interface{}, err error) {
	w.wheres = append(w.wheres, WhereExpr{
		oper: AndOper,
		cond: "1=1",
	})

	var seg string
	for _, v := range w.wheres {
		seg, _, err = v.ToSQL()
		if err != nil {
			return
		}
		sql += seg
	}

	return
}
