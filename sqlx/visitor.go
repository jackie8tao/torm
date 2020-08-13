package sqlx

// Visitor visits the expression or statments to generate sql string.
type Visitor interface {
	Visit(Stmt) error
}
