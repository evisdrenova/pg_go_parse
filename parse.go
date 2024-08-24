package pggoquery

import (
	"github.com/evisdrenova/pg_go_parse/internal/node"
	"github.com/wasilibs/go-pgquery/parser"
	"google.golang.org/protobuf/proto"
)

type ParseResult struct {
	Version int32
	Stmts   []*RawStmt
}

type RawStmt struct {
	Stmt         *node.Node
	StmtLocation int32
	StmtLen      int32
}

// Parse the given SQL statement into a parse tree (Go struct format)
func Parse(input string) (tree *ParseResult, err error) {
	protobufTree, err := parser.ParseToProtobuf(input)
	if err != nil {
		return
	}

	tree = &ParseResult{}
	err = proto.Unmarshal(protobufTree, tree)
	return
}
