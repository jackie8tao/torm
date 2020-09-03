package sqlgen

import (
	"fmt"
	"strings"

	"git.code.oa.com/pluto/torm/expr"
)

// SelectStmt select statement.
type SelectStmt struct {
	distinct bool
	columns  expr.Expr
	from     expr.Expr
	joins    expr.Expr
	wheres   expr.Expr
	groupBys expr.Expr
	havings  expr.Expr
	ordersBy expr.Expr
	limit    int
	offset   int
}

// selectInjector writes 'SELECT' string to the select statement.
func (s *SelectStmt) selectInjector(builder SQLBuilder) error {
	return builder.WriteString("SELECT")
}

// distinctInjector injects distinct segment to the select statement if true.
func (s *SelectStmt) distinctInjector(builder SQLBuilder) error {
	if s.distinct {
		err := builder.WriteString(" DISTINCT")
		if err != nil {
			return err
		}
	}
	return nil
}

// colListInjector injects columns segment to the select statement.
func (s *SelectStmt) colListInjector(builder SQLBuilder) error {
	if s.columns == nil {
		return nil
	}

	return builder.WriteExpr(s.columns)
}

// fromInjector injects from segment to the select statement.
func (s *SelectStmt) fromInjector(builder SQLBuilder) error {
	if s.from == nil {
		return nil
	}

	err := builder.WriteString(" FROM")
	if err != nil {
		return err
	}

	return builder.WriteExpr(s.from)
}

// joinsInjector injects join segment to the select statement.
func (s *SelectStmt) joinsInjector(builder SQLBuilder) error {
	if s.joins == nil {
		return nil
	}

	return builder.WriteExpr(s.joins)
}

// wheresInjector injects where segment to the select statement.
func (s *SelectStmt) wheresInjector(builder SQLBuilder) error {
	if s.wheres == nil {
		return nil
	}

	err := builder.WriteString(" WHERE")
	if err != nil {
		return err
	}

	return builder.WriteExpr(s.wheres)
}

// groupBysInjector injects group-by segment to the select statement.
func (s *SelectStmt) groupBysInjector(builder SQLBuilder) error {
	if s.groupBys == nil {
		return nil
	}

	err := builder.WriteString(" GROUP BY")
	if err != nil {
		return err
	}

	return builder.WriteExpr(s.groupBys)
}

// havingsInjector injects having segment into the select statement.
func (s *SelectStmt) havingsInjector(builder SQLBuilder) error {
	if s.havings == nil {
		return nil
	}

	err := builder.WriteString(" HAVING")
	if err != nil {
		return err
	}

	return builder.WriteExpr(s.havings)
}

// ordersByInjector injects order-by segment into the select statement.
func (s *SelectStmt) ordersByInjector(builder SQLBuilder) error {
	if s.ordersBy == nil {
		return nil
	}

	err := builder.WriteString(" ORDER BY")
	if err != nil {
		return err
	}

	return builder.WriteExpr(s.ordersBy)
}

// limitInjector injects limit segment into the select statement.
func (s *SelectStmt) limitInjector(builder SQLBuilder) error {
	if s.limit <= 0 {
		return nil
	}

	return builder.WriteString(fmt.Sprintf(" LIMIT %d", s.limit))
}

// offsetInjector injects offset segment into the select statement.
func (s *SelectStmt) offsetInjector(builder SQLBuilder) error {
	if s.offset <= 0 {
		return nil
	}

	return builder.WriteString(fmt.Sprintf(" OFFSET %d", s.offset))
}

// ToSQL returns the select sql and arguments.if having a error, it will return error.
func (s *SelectStmt) ToSQL() (sql string, args []interface{}, err error) {
	builder := &bufSQLBuilder{
		buf:  &strings.Builder{},
		args: []interface{}{},
	}

	injectors := []injector{
		s.selectInjector,
		s.distinctInjector,
		s.colListInjector,
		s.fromInjector,
		s.joinsInjector,
		s.wheresInjector,
		s.groupBysInjector,
		s.havingsInjector,
		s.ordersByInjector,
		s.limitInjector,
		s.offsetInjector,
	}
	for _, v := range injectors {
		if err = v(builder); err != nil {
			return
		}
	}

	return builder.ToSQL()
}
