package torm

import (
	"bytes"
)

// SelectBuilder
type SelectBuilder struct {
	distinct    bool
	columns     []Sqlizer
	fromParts   []Sqlizer
	joins       []Sqlizer
	whereParts  []Sqlizer
	groupBys    []string
	havingParts []Sqlizer
	orderBys    []string
	limit       uint64
	offset      uint64
}

func (bd *SelectBuilder) distinctAssign(buf *bytes.Buffer) (err error) {
	if bd.distinct {
		_, err = buf.WriteString("")
	}
	return
}

// ToSQL
func (bd *SelectBuilder) ToSQL() (sql string, args []interface{}, err error) {
	assigns :=
}
