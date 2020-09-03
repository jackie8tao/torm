package expr

import (
	"fmt"
	"strings"
)

// ColExpr column expression.
type ColExpr struct {
	table  string
	column string
}

// ToSQL this function implements Expr interface.
func (e ColExpr) ToSQL() (sql string, args []interface{}, err error) {
	if e.table != "" {
		sql += fmt.Sprintf("`%s`.", e.table)
	}
	if e.column != "" {
		sql += fmt.Sprintf(" `%s`", e.column)
		return
	}
	sql += " `*`"
	return
}

// ColList column expression list.
type ColList struct {
	cols []ColExpr
}

// ToSQL this function implements Expr interface.
func (e ColList) ToSQL() (sql string, args []interface{}, err error) {
	var part string
	for _, v := range e.cols {
		part, _, err = v.ToSQL()
		if err != nil {
			return
		}
		sql += fmt.Sprintf("%s,", part)
	}
	sql = strings.TrimSuffix(sql, ",")
	return
}

// ValueExpr value expression used in update statement.
type ValueExpr struct {
	table string
	col   string
	arg   interface{}
}

// ToSQL this function implements Expr interface.
func (v ValueExpr) ToSQL() (sql string, args []interface{}, err error) {
	args = make([]interface{}, 0)
	sql += " "
	if v.table != "" {
		sql += fmt.Sprintf("`%s`.", v.table)
	}
	if v.col == "" {
		err = ErrEmptyCol
		return
	}
	sql += fmt.Sprintf("`%s` = ?", v.col)
	args = append(args, v.arg)
	return
}

// ValueList value expression list.
type ValueList struct {
	cols []ValueExpr
}

// ToSQL this function implements Expr interface.
func (v ValueList) ToSQL() (sql string, args []interface{}, err error) {
	var (
		seg string
		arg []interface{}
	)

	args = make([]interface{}, 0)

	for _, v := range v.cols {
		seg, arg, err = v.ToSQL()
		if err != nil {
			return
		}
		args = append(args, arg...)
		sql += fmt.Sprintf(" %s,", seg)
	}

	sql = strings.TrimSuffix(sql, ",")
	return
}
