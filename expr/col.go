package expr

import (
	"fmt"
	"strings"
)

// ColExpr column expression
type ColExpr struct {
	table  string
	column string
}

// ToSQL this function implements sqlizer interface.
func (e ColExpr) ToSQL() (sql string, args []interface{}, err error) {
	if e.table != "" {
		sql += fmt.Sprintf("`%s`.", e.table)
	}
	sql += fmt.Sprintf("`%s`", e.column)
	return
}

// ColList column expression list.
type ColList struct {
	cols []ColExpr
}

// ToSQL
func (e ColList) ToSQL() (sql string, args []interface{}, err error) {
	var part string
	for _, v := range e.cols {
		part, _, err = v.ToSQL()
		if err != nil {
			return
		}
		sql += fmt.Sprintf(" %s,", part)
	}
	sql = strings.TrimSuffix(sql, ",")
	return
}
