package torm

import (
	"fmt"
)

// Sqlizer is the interface that wraps the ToSql method.
type Sqlizer interface {
	// ToSql returns a SQL representation of the Sqlizer, along with a slice of args
	// as passed to e.g. database/sql.Exec. It can also return an error.
	ToSQL() (string, []interface{}, error)
}

// SqlSetter
type SqlSetter func(SQLBuilder) error

// ColExpr
type ColExpr struct {
	table  string
	column string
}

// ToSQL this function implements sqlizer interface.
func (e ColExpr) ToSQL() (sql string, args []interface{}, err error) {
	if e.table != "" {
		sql += fmt.Sprintf("`%s`.", e.table)
	}
	sql += fmt.Sprintf("`%s`, ", e.column)
	return
}

// FromExpr
type FromExpr struct {
	table string
}

// ToSQL
func (e FromExpr) ToSQL() (sql string, args []interface{}, err error) {
	sql += fmt.Sprintf("`%s`", e.table)
	return
}

const (
	LeftJoin = iota + 1
	RightJoin
)

// JoinExpr
type JoinExpr struct {
	method int
	table  string
	cond   string
}

func (e JoinExpr) ToSQL() (sql string, args []interface{}, err error) {
	var method string
	switch e.method {
	case LeftJoin:
		method = "LEFT JOIN"
	case RightJoin:
		method = "RIGHT JOIN"
	}
	sql = fmt.Sprintf("%s %s ON %s", method, e.table, e.cond)
	return
}

const (
	AndOper = iota + 1
	OrOper
)

// WhereExpr
type WhereExpr struct {
	oper int
	cond string
}

func (e WhereExpr) ToSQL() (sql string, args []interface{}, err error) {
	var oper string
	switch e.oper {
	case AndOper:
		oper += " AND "
	case OrOper:
		oper += " OR "
	default:
	}

	sql = fmt.Sprintf("%s (%s)", oper, e.cond)
	return
}

// GroupByExpr
type GroupByExpr struct {
	col string
}

func (e GroupByExpr) ToSQL() (sql string, args []interface{}, err error) {
	sql += fmt.Sprintf("`%s`", e.col)
	return
}

// OrderByExpr
type OrderByExpr struct {
	col  string
	desc bool
}

func (e OrderByExpr) ToSQL() (sql string, args []interface{}, err error) {
	sql += fmt.Sprintf("`%s`", e.col)
	if e.desc {
		sql += " DESC"
	} else {
		sql += " ASC"
	}
	return
}
