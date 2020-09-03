package expr

import (
	"fmt"
)

const (
	LeftJoin JoinMethod = iota + 1
	RightJoin
	InnerJoin
)

var supportedMethods = map[JoinMethod]string{
	LeftJoin:  "LEFT JOIN",
	RightJoin: "RIGHT JOIN",
	InnerJoin: "INNER JOIN",
}

// JoinMethod join type, eg: LEFT JOIN, RIGHT JOIN and so on.
type JoinMethod int

// JoinExpr join expression.
type JoinExpr struct {
	method JoinMethod
	table  string
	cond   string
}

func (j JoinExpr) convertMethod(method JoinMethod) (val string, err error) {
	var ok bool
	val, ok = supportedMethods[method]
	if !ok {
		err = ErrInvalidMethod
	}
	return
}

// ToSQL this function implements Expr interface.
func (j JoinExpr) ToSQL() (sql string, args []interface{}, err error) {
	method, err := j.convertMethod(j.method)
	if err != nil {
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
