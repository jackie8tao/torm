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
func (bsb *bufSQLBuilder) WriteString(str string) error {
	_, err := bsb.buf.WriteString(str)
	return err
}

// WriteSqlizer
func (bsb *bufSQLBuilder) WriteSqlizer(sqlizer Sqlizer) error {
	sql, _, err := sqlizer.ToSQL()
	if err != nil {
		return err
	}
	return bsb.WriteString(sql)
}

// ToSQL
func (bsb *bufSQLBuilder) ToSQL() string {
	return bsb.buf.String()
}
