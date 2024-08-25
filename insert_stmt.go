package pggoquery

type InsertStmt struct {
	TableName string
	Columns   []string
	Values    [][]*Node
}
