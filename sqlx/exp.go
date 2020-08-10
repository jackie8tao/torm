package sqlx

// Exp
type Exp interface {
	Visit(SQLBuilder) error
}
