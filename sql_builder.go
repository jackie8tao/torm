package torm

import (
	"bytes"
)

type (
	// SQLBuilder
	SQLBuilder interface {
		WriteString(string)
	}

	// bufSQLBuilder
	bufSQLBuilder struct {
		buf *bytes.Buffer
	}
)

func (bsb *bufSQLBuilder) WriteString(str string) {
	_, err := bsb.buf.WriteString(str)
	if err != nil {
		panic(err)
	}
}
