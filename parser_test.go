package pggoquery

import (
	"testing"
)

func Test_ParseStatementSelectBasic(t *testing.T) {
	input := `SELECT name, age FROM users WHERE age > 20`

	l := NewLexer(input)
	p := newParser(l)

	selectStmt, err := p.ParseStatement()
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	stmt, ok := selectStmt.NodeType.(*SelectStmt)
	if !ok {
		t.Fatalf("expected *SelectStmt, got %T", stmt)
	}

	if len(stmt.TargetList) != 2 {
		t.Fatalf("stmt.TargetList has wrong length. got=%d", len(stmt.TargetList))
	}

	if len(stmt.FromClause) != 1 {
		t.Fatalf("stmt.FromClause has wrong length. got=%d", len(stmt.FromClause))
	}

	if stmt.WhereClause == nil {
		t.Fatal("stmt.WhereClause is nil")
	}
}

func Test_ParseStatementSelectStar(t *testing.T) {
	input := `SELECT * FROM users WHERE age > 20`

	l := NewLexer(input)
	p := newParser(l)

	selectStmt, err := p.ParseStatement()
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	stmt, ok := selectStmt.NodeType.(*SelectStmt)
	if !ok {
		t.Fatalf("expected *SelectStmt, got %T", stmt)
	}

	if len(stmt.TargetList) != 1 {
		t.Fatalf("stmt.TargetList has wrong length. got=%d", len(stmt.TargetList))
	}

	if len(stmt.FromClause) != 1 {
		t.Fatalf("stmt.FromClause has wrong length. got=%d", len(stmt.FromClause))
	}

	if stmt.WhereClause == nil {
		t.Fatal("stmt.WhereClause is nil")
	}
}

func Test_ParseStatementSelectManyIdentifiers(t *testing.T) {
	input := `SELECT name, age, first, last, middle, street, city, state FROM users WHERE age > 20`

	l := NewLexer(input)
	p := newParser(l)

	selectStmt, err := p.ParseStatement()
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	stmt, ok := selectStmt.NodeType.(*SelectStmt)
	if !ok {
		t.Fatalf("expected *SelectStmt, got %T", stmt)
	}

	if len(stmt.TargetList) != 8 {
		t.Fatalf("stmt.TargetList has wrong length. got=%d", len(stmt.TargetList))
	}

	if len(stmt.FromClause) != 1 {
		t.Fatalf("stmt.FromClause has wrong length. got=%d", len(stmt.FromClause))
	}

	if stmt.WhereClause == nil {
		t.Fatal("stmt.WhereClause is nil")
	}
}

// func Test_ParseStatementSelectFullyQualifiedColumns(t *testing.T) {
// 	input := `SELECT name, age, first FROM users WHERE age > 20`

// 	l := NewLexer(input)
// 	p := newParser(l)

// 	selectStmt, err := p.ParseStatement()
// 	if err != nil {
// 		t.Fatalf("parser error: %v", err)
// 	}

// 	stmt, ok := selectStmt.NodeType.(*SelectStmt)
// 	if !ok {
// 		t.Fatalf("expected *SelectStmt, got %T", stmt)
// 	}

// 	if len(stmt.TargetList) != 8 {
// 		t.Fatalf("stmt.TargetList has wrong length. got=%d", len(stmt.TargetList))
// 	}

// 	if len(stmt.FromClause) != 1 {
// 		t.Fatalf("stmt.FromClause has wrong length. got=%d", len(stmt.FromClause))
// 	}

// 	if stmt.WhereClause == nil {
// 		t.Fatal("stmt.WhereClause is nil")
// 	}
// }

func Test_parseTargetList(t *testing.T) {

	input := "name, age, address"

	l := NewLexer(input)
	p := newParser(l)

	targetList := p.parseTargetList()

	expectedLiterals := []string{"name", "age", "address"}

	if len(targetList) != len(expectedLiterals) {
		t.Fatalf("expected %d nodes, got %d", len(expectedLiterals), len(targetList))
	}

	for i, node := range targetList {
		if node.NodeType != expectedLiterals[i] {
			t.Errorf("expected node %d to be %s, got %s", i, expectedLiterals[i], node.NodeType)
		}
	}
}

func Test_parseFromClause(t *testing.T) {
	input := "table.users"

	l := NewLexer(input)
	p := newParser(l)

	fromClause := p.parseFromClause()

	expectedLiterals := []string{"table", "users"}

	if len(fromClause) != len(expectedLiterals) {
		t.Fatalf("expected %d nodes, got %d", len(expectedLiterals), len(fromClause))
	}

	for i, node := range fromClause {
		if node.NodeType != expectedLiterals[i] {
			t.Errorf("expected node %d to be %s, got %s", i, expectedLiterals[i], node.NodeType)
		}
	}
}
