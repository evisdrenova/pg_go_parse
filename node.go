package pggoquery

type Node struct {
	NodeType interface{} // defines the type of node i.e statement, for full list check node_types.text
}

type NodeSelectStmt struct {
	SelectStmt *NodeSelectStmt
}
