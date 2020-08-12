package sqlx

// expression type values
const (
	IdentifierExp ExpType = iota + 1
)

// ExpType expression type
type ExpType int

// Exp sql expression
type Exp interface {
	Type() ExpType
	Visit(SQLBuilder) error
}
