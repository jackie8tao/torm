package torm

import (
	"strings"
)

// SelectStmt select语句
type SelectStmt struct {
	Root Node
}

// ToSQL 返回sql语句
func (s *SelectStmt) ToSQL() (sql string, err error) {
	sb := &strings.Builder{}

	err = s.Root.Walk(sb)
	if err != nil {
		return
	}

	sql = sb.String()
	return
}
