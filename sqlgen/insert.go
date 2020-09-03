package sqlgen

// InsertStmt insert sql statement.
type InsertStmt struct {
	Table        string
	Column       []string
	Value        [][]interface{}
	Ignored      bool
	ReturnColumn []string
	RecordID     *int64
}
