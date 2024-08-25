package pggoquery

type SelectStmt struct {
	DistinctClause []*Node
	IntoClause     *IntoClause
	TargetList     []*Node
	FromClause     []*Node
	WhereClause    *Node
	GroupClause    []*Node
	GroupDistinct  bool
	HavingClause   *Node
	WindowClause   []*Node
	ValuesLists    []*Node
	SortClause     []*Node
	LimitOffset    *Node
	LimitCount     *Node
	LimitOption    LimitOption
	LockingClause  []*Node
	WithClause     *WithClause
	Op             SetOperation
	All            bool
	Larg           *SelectStmt
	Rarg           *SelectStmt
}

type IntoClause struct {
	Rel            *RangeVar
	ColNames       []*Node
	AccessMethod   string
	Options        []*Node
	OnCommit       OnCommitAction
	TableSpaceName string
	ViewQuery      *Node
	SkipData       bool
}

type RangeVar struct {
	Catalogname    string
	Schemaname     string
	Relname        string
	Inh            bool
	Relpersistence string
	Alias          *Alias
	Location       int32
}

type Alias struct {
	Aliasname string
	Colnames  []*Node
}

type LimitOption int32

const (
	LimitOption_LIMIT_OPTION_UNDEFINED LimitOption = 0
	LimitOption_LIMIT_OPTION_DEFAULT   LimitOption = 1
	LimitOption_LIMIT_OPTION_COUNT     LimitOption = 2
	LimitOption_LIMIT_OPTION_WITH_TIES LimitOption = 3
)

// Enum value maps for LimitOption.
var (
	LimitOption_name = map[int32]string{
		0: "LIMIT_OPTION_UNDEFINED",
		1: "LIMIT_OPTION_DEFAULT",
		2: "LIMIT_OPTION_COUNT",
		3: "LIMIT_OPTION_WITH_TIES",
	}
	LimitOption_value = map[string]int32{
		"LIMIT_OPTION_UNDEFINED": 0,
		"LIMIT_OPTION_DEFAULT":   1,
		"LIMIT_OPTION_COUNT":     2,
		"LIMIT_OPTION_WITH_TIES": 3,
	}
)

type WithClause struct {
	Ctes      []*Node
	Recursive bool
	Location  int32
}

type SetOperation int32

const (
	SetOperation_SET_OPERATION_UNDEFINED SetOperation = 0
	SetOperation_SETOP_NONE              SetOperation = 1
	SetOperation_SETOP_UNION             SetOperation = 2
	SetOperation_SETOP_INTERSECT         SetOperation = 3
	SetOperation_SETOP_EXCEPT            SetOperation = 4
)

// Enum value maps for SetOperation.
var (
	SetOperation_name = map[int32]string{
		0: "SET_OPERATION_UNDEFINED",
		1: "SETOP_NONE",
		2: "SETOP_UNION",
		3: "SETOP_INTERSECT",
		4: "SETOP_EXCEPT",
	}
	SetOperation_value = map[string]int32{
		"SET_OPERATION_UNDEFINED": 0,
		"SETOP_NONE":              1,
		"SETOP_UNION":             2,
		"SETOP_INTERSECT":         3,
		"SETOP_EXCEPT":            4,
	}
)

type OnCommitAction int32

const (
	OnCommitAction_ON_COMMIT_ACTION_UNDEFINED OnCommitAction = 0
	OnCommitAction_ONCOMMIT_NOOP              OnCommitAction = 1
	OnCommitAction_ONCOMMIT_PRESERVE_ROWS     OnCommitAction = 2
	OnCommitAction_ONCOMMIT_DELETE_ROWS       OnCommitAction = 3
	OnCommitAction_ONCOMMIT_DROP              OnCommitAction = 4
)

// Enum value maps for OnCommitAction.
var (
	OnCommitAction_name = map[int32]string{
		0: "ON_COMMIT_ACTION_UNDEFINED",
		1: "ONCOMMIT_NOOP",
		2: "ONCOMMIT_PRESERVE_ROWS",
		3: "ONCOMMIT_DELETE_ROWS",
		4: "ONCOMMIT_DROP",
	}
	OnCommitAction_value = map[string]int32{
		"ON_COMMIT_ACTION_UNDEFINED": 0,
		"ONCOMMIT_NOOP":              1,
		"ONCOMMIT_PRESERVE_ROWS":     2,
		"ONCOMMIT_DELETE_ROWS":       3,
		"ONCOMMIT_DROP":              4,
	}
)
