package torm

import (
	"bytes"
)

type (
	// SQLBuilder
	SQLBuilder interface {
		WriteString(string) error
		WriteSqlizer(Sqlizer) error
		ToSQL() string
	}

	// bufSQLBuilder
	bufSQLBuilder struct {
		buf *bytes.Buffer
	}
)

// WriteString
func (b *bufSQLBuilder) WriteString(str string) error {
	_, err := b.buf.WriteString(str)
	return err
}

// WriteSqlizer
func (b *bufSQLBuilder) WriteSqlizer(sqlizer Sqlizer) error {
	sql, _, err := sqlizer.ToSQL()
	if err != nil {
		return err
	}
	return b.WriteString(sql)
}

// ToSQL
func (b *bufSQLBuilder) ToSQL() string {
	return b.buf.String()
}
