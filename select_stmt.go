package torm

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
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
	limit       int
	offset      int
}

// distinctSetter
func (sb *SelectBuilder) distinctSetter(bd SQLBuilder) (err error) {
	if sb.distinct {
		err = bd.WriteString(" DISTINCT ")
	}
	return
}

// columnSetter
func (sb *SelectBuilder) columnSetter(bd SQLBuilder) (err error) {
	if len(sb.columns) <= 0 {
		return
	}

	for _, v := range sb.columns {
		err = bd.WriteSqlizer(v)
		if err != nil {
			return
		}
	}

	return
}

// fromSetter
func (sb *SelectBuilder) fromSetter(bd SQLBuilder) (err error) {
	if len(sb.fromParts) <= 0 {
		err = errors.New("illegal where part")
		return
	}

	err = bd.WriteString("FROM ")
	if err != nil {
		return
	}

	for _, v := range sb.fromParts {
		err = bd.WriteSqlizer(v)
		if err != nil {
			return
		}
	}

	return
}

// joinSetter
func (sb *SelectBuilder) joinSetter(bd SQLBuilder) (err error) {
	if len(sb.joins) <= 0 {
		return
	}

	for _, v := range sb.joins {
		err = bd.WriteSqlizer(v)
		if err != nil {
			return
		}
	}

	return
}

// whereSetter
func (sb *SelectBuilder) whereSetter(bd SQLBuilder) (err error) {
	if len(sb.whereParts) <= 0 {
		return
	}

	err = bd.WriteString("WHERE ")
	if err != nil {
		return
	}

	for _, v := range sb.whereParts {
		err = bd.WriteSqlizer(v)
		if err != nil {
			return
		}
	}

	return
}

// groupBySetter
func (sb *SelectBuilder) groupBySetter(bd SQLBuilder) (err error) {
	if len(sb.groupBys) <= 0 {
		return
	}
	err = bd.WriteString("GROUP BY")
	if err != nil {
		return
	}
	return bd.WriteString(strings.Join(sb.groupBys, " "))
}

// havingSetter
func (sb *SelectBuilder) havingSetter(bd SQLBuilder) (err error) {
	if len(sb.havingParts) <= 0 {
		return
	}

	err = bd.WriteString("HAVING")
	if err != nil {
		return
	}

	for _, v := range sb.havingParts {
		err = bd.WriteSqlizer(v)
		if err != nil {
			return
		}
	}

	return
}

// orderBySetter
func (sb *SelectBuilder) orderBySetter(bd SQLBuilder) (err error) {
	if len(sb.orderBys) <= 0 {
		return
	}
	err = bd.WriteString("ORDER BY ")
	if err != nil {
		return
	}
	return bd.WriteString(strings.Join(sb.orderBys, " "))
}

// limitSetter
func (sb *SelectBuilder) limitSetter(bd SQLBuilder) (err error) {
	if sb.limit <= 0 {
		return
	}
	err = bd.WriteString("LIMIT ")
	if err != nil {
		return
	}
	return bd.WriteString(strconv.Itoa(sb.limit))
}

// offsetSetter
func (sb *SelectBuilder) offsetSetter(bd SQLBuilder) (err error) {
	if sb.offset <= 0 {
		return
	}
	err = bd.WriteString("OFFSET ")
	if err != nil {
		return
	}
	return bd.WriteString(strconv.Itoa(sb.offset))
}

// ToSQL
func (sb *SelectBuilder) ToSQL() (sql string, args []interface{}, err error) {
	bd := &bufSQLBuilder{
		buf: &bytes.Buffer{},
	}

	err = bd.WriteString("SELECT ")
	if err != nil {
		return
	}

	setters := []SqlSetter{
		sb.distinctSetter, sb.columnSetter, sb.fromSetter,
		sb.whereSetter, sb.groupBySetter, sb.havingSetter,
		sb.orderBySetter, sb.limitSetter, sb.offsetSetter,
	}
	for _, v := range setters {
		err = v(bd)
		if err != nil {
			return
		}
	}

	sql = bd.ToSQL()
	return
}
