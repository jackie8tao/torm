package stmt

import "github.com/jackie8tao/torm/sqlx"

type (
	// AliasStmt
	AliasStmt interface {
		sqlx.Stmt
	}

	aliasStmt struct {
		raw     identStmt
		aliased sqlx.Stmt
	}
)

func (a aliasStmt) Type() sqlx.StmtType {
	return sqlx.AliasStmt
}

func (a aliasStmt) Stmt() sqlx.Stmt {
	return a
}

func (a aliasStmt) Clone() sqlx.Stmt {
	return aliasStmt{
		raw:     a.raw,
		aliased: a.aliased,
	}
}
