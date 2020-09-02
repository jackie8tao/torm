package sqlgen

import (
	"strings"

	"git.code.oa.com/pluto/torm/expr"
)

// SelectStmt select statement.
type SelectStmt struct {
	distinct bool
	columns  expr.ColList
}

func (s *SelectStmt) selectInjector(builder SQLBuilder) error {
	return builder.WriteString("SELECT")
}

func (s *SelectStmt) distinctInjector(builder SQLBuilder) error {
	if s.distinct {
		err := builder.WriteString(" DISTINCT")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SelectStmt) colListInjector(builder SQLBuilder) error {
	return injectExpr(builder, s.columns)
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
	}
	for _, v := range injectors {
		if err = v(builder); err != nil {
			return
		}
	}

	return builder.ToSQL()
}
