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
	ErrEmptyTable    = errors.New("empty table value")
	ErrInvalidMethod = errors.New("invalid join method")
	ErrInvalidOper   = errors.New("invalid where operator")
	ErrInvalidOrder  = errors.New("invalid order-by value")
	ErrEmptyCol      = errors.New("empty column value")
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

// NewTableExpr creates from expression.
func NewTableExpr(table string) TableExpr {
	return TableExpr{
		table: table,
	}
}

// NewWhereExpr creates where expression.
func NewWhereExpr(oper OperVal, cond string, args ...interface{}) WhereExpr {
	return WhereExpr{
		oper: oper,
		cond: cond,
		args: args,
	}
}

// NewWhereList creates where expression list.
func NewWhereList(exprs ...WhereExpr) WhereList {
	return WhereList{
		wheres: exprs,
	}
}

// NewJoinExpr creates join expression.
func NewJoinExpr(method JoinMethod, table, cond string) JoinExpr {
	return JoinExpr{
		method: method,
		table:  table,
		cond:   cond,
	}
}

// NewJoinList creates join expression list.
func NewJoinList(joins ...JoinExpr) JoinList {
	return JoinList{
		joins: joins,
	}
}

// NewOrderByExpr creates order-by expression.
func NewOrderByExpr(order OrderVal, col string) OrderByExpr {
	return OrderByExpr{
		order: order,
		col:   col,
	}
}

// NewOrderByList creates order-by expression list.
func NewOrderByList(orders ...OrderByExpr) OrderByList {
	return OrderByList{
		orders: orders,
	}
}

// NewValueExpr creates value expression which is used in update statement.
func NewValueExpr(col string, val interface{}) ValueExpr {
	detail := strings.Split(col, ".")
	if len(detail) >= 2 {
		return ValueExpr{
			table: detail[0],
			col:   detail[1],
			arg:   val,
		}
	}
	return ValueExpr{
		col: col,
		arg: val,
	}
}

// NewValueList creates value expression list.
func NewValueList(vals ...ValueExpr) ValueList {
	return ValueList{
		cols: vals,
	}
}
