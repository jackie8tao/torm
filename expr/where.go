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
	sql = fmt.Sprintf(" (%s)", w.cond)
	return
}

// WhereList where list expression.
type WhereList struct {
	wheres []WhereExpr
}

func (w WhereList) convertOper(oper OperVal) (val string, err error) {
	switch oper {
	case OrOper:
		val = "OR"
	case AndOper:
		val = "AND"
	default:
		err = ErrInvalidOper
	}
	return
}

// ToSQL this function implements Expr interface.
func (w WhereList) ToSQL() (sql string, args []interface{}, err error) {
	var (
		seg  string
		oper string
	)

	for k, v := range w.wheres {
		if k == 0 {
			seg = fmt.Sprintf(" (%s)", v.cond)
		} else {
			oper, err = w.convertOper(v.oper)
			if err != nil {
				return
			}
			seg = fmt.Sprintf(" %s (%s)", oper, v.cond)
		}
		sql += seg
	}

	return
}
