package stmt

import "github.com/jackie8tao/torm/sqlx"

// IdentStmt interface for identStmt expression.
type IdentStmt interface {
	sqlx.Aliaseable
	sqlx.Comparable
	sqlx.Inable
	sqlx.Isable
	sqlx.Likeable
	sqlx.Orderable
	sqlx.Distinctable
	sqlx.Castable
	// returns true if this identifier has more more than on part (Schema, Table or Col)
	//	"schema" -> true //cant qualify anymore
	//	"schema.table" -> true
	//	"table" -> false
	// "schema"."table"."col" -> true
	// "table"."col" -> true
	// "col" -> false
	IsQualified() bool
	// Returns a new IdentStmt with the specified schema
	Schema(string) IdentStmt
	// Returns the current schema
	GetSchema() string
	// Returns a new IdentStmt with the specified table
	Table(string) IdentStmt
	// Returns the current table
	GetTable() string
	// Returns a new IdentStmt with the specified column
	Col(interface{}) IdentStmt
	// Returns the current column
	GetCol() interface{}
	// Returns a new IdentifierExpression with the column set to *
	//   I("my_table").All() //"my_table".*
	All() IdentStmt

	// Returns true if schema table and identifier are all zero values.
	IsEmpty() bool
}

// sql expression for identStmt
type identStmt struct {
	schema string
	table  string
	col    interface{}
}

// NewIdentExp
func NewIdentExp(schema, table string, col interface{}) sqlx.Stmt {
	return identStmt{
		schema: schema,
		table:  table,
		col:    col,
	}
}

// Type get the type value of identStmt.
func (i identStmt) Type() sqlx.StmtType {
	return sqlx.IdentStmt
}

// Clone clone the current expression
func (i identStmt) Clone() sqlx.Stmt {
	return identStmt{
		schema: i.schema,
		table:  i.table,
		col:    i.col,
	}
}

// Expr returns the current expression.
func (i identStmt) Stmt() sqlx.Stmt {
	return i
}

// TODO check qualified
func (i identStmt) IsQualified() bool {
	return false
}

// // Table sets the table on the current identifier expression.
// func (i identStmt) Table(table string) IdentStmt {
// 	i.table = table
// 	return i
// }
