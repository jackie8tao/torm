package expr

import (
	"errors"
	"strings"
)

// Expr is the interface that all sql segments should implements.
type Expr interface {
	// ToSQL returns a SQL representation of the expression, along with a slice of args
	// as passed to e.g. database/sql.Exec. It can also return an error.
	ToSQL() (string, []interface{}, error)
}

// expression errors.
var (
	ErrEmptyTable           = errors.New("empty table value")
	ErrIllegalJoinMethod    = errors.New("illegal join method")
	ErrIllegalWhereOperator = errors.New("illegal where operator")
	ErrIllegalOrderBy       = errors.New("illegal order-by value")
)

// NewColExpr creates column expression.
func NewColExpr(col string) ColExpr {
	detail := strings.Split(col, ".")
	switch len(detail) {
	case 1:
		return ColExpr{
			column: col,
		}
	default:
		return ColExpr{
			table:  detail[0],
			column: detail[1],
		}
	}
}

// NewColList creates column expression list.
func NewColList(cols ...string) ColList {
	list := make([]ColExpr, 0, len(cols))
	for _, v := range cols {
		list = append(list, NewColExpr(v))
	}

	return ColList{
		cols: list,
	}
}

// NewFromExpr creates from expression.
func NewFromExpr(table string) FromExpr {
	return FromExpr{
		table: table,
	}
}
