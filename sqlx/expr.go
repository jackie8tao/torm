package sqlx

import (
	"bytes"
	"fmt"
)

type expr struct {
	sql  string
	args []interface{}
}

func Expr(sql string, args ...interface{}) Sqlizer {
	return expr{
		sql:  sql,
		args: args,
	}
}

func (e expr) ToSql() (string, []interface{}, error) {
	if !hasSqlizer(e.args) {
		return e.sql, e.args, nil
	}

	args := make([]interface{}, 0, len(e.args))
	sql, err := replacePlaceholders(e.sql, func(buf *bytes.Buffer, i int) error {
		if i > len(e.args) {
			buf.WriteRune('?')
			return nil
		}
		switch arg := e.args[i-1].(type) {
		case Sqlizer:
			sql, vs, err := arg.ToSql()
			if err != nil {
				return err
			}
			args = append(args, vs...)
			fmt.Fprintf(buf, sql)
		default:
			args = append(args, arg)
			buf.WriteRune('?')
		}
		return nil
	})
	if err != nil {
		return "", nil, err
	}
	return sql, args, nil
}

func hasSqlizer(args []interface{}) bool {
	for _, arg := range args {
		_, ok := arg.(Sqlizer)
		if ok {
			return true
		}
	}
	return false
}
