package expr

import (
	"fmt"
	"strings"
)

// order values.
const (
	AscOrder OrderVal = iota + 1
	DescOrder
)

// OrderVal order types.
type OrderVal int

// OrderByExpr order-by expression.
type OrderByExpr struct {
	order OrderVal
	col   string
}

// ToSQL this function implements Expr interface.
func (o OrderByExpr) ToSQL() (sql string, args []interface{}, err error) {
	var order string
	switch o.order {
	case AscOrder:
		order = "ASC"
	case DescOrder:
		order = "DESC"
	default:
		err = ErrIllegalOrderBy
		return
	}

	sql = fmt.Sprintf(" %s %s,", o.col, order)

	return
}

// OrderByList order-by list expression.
type OrderByList struct {
	orders []OrderByExpr
}

// ToSQL this function implement Expr interface.
func (o OrderByList) ToSQL() (sql string, args []interface{}, err error) {
	var seg string
	for _, v := range o.orders {
		seg, _, err = v.ToSQL()
		if err != nil {
			return
		}
		sql += fmt.Sprintf("%s,", seg)
	}

	sql = strings.TrimSuffix(sql, ",")

	return
}
