package pggoquery

import "fmt"

// defines a parser that can create an abstract syntax tree (AST)
// ultimatlty, this should patch the underlying AST that postgres creates
type Parser struct {
	l       *Lexer
	curTok  Token
	peekTok Token
}

// create a new Parser that can build an AST from the lexer
func newParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken() // read two tokens so the curTok and peekTok are both set
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

// parses select statements
func (p *Parser) ParseStatement() (*Node, error) {

	switch p.curTok.Literal {
	case "SELECT":
		return p.ParseSelectStatement()
	// case "INSERT":
	// 	return p.ParseInsertStatement()
	default:
		return nil, fmt.Errorf("unexpected statement: %s", p.curTok.Literal)
	}
}

func (p *Parser) ParseSelectStatement() (*Node, error) {

	stmt := &SelectStmt{}

	// if the next token is not a keyword or it's not SELECT then return
	if p.curTok.Literal != "SELECT" {
		return nil, fmt.Errorf("expected SELECT, got %s", p.curTok.Literal)
	}

	// go to next token
	p.nextToken()

	// parse the target list for the select which is a []*Node
	stmt.TargetList = p.parseTargetList()

	// if the next token is not not FROM then return
	if p.curTok.Literal != "FROM" {
		return nil, fmt.Errorf("expected FROM, got %s", p.curTok.Literal)
	}
	// go to next token
	p.nextToken()

	// parse the from clause for the select which is a []*Node
	stmt.FromClause = p.parseFromClause()

	// parse the where clause for the select which is a []*Node
	if p.curTok.Literal == "WHERE" {
		p.nextToken()
		stmt.WhereClause = p.parseWhereClause()
	}

	return &Node{NodeType: stmt}, nil
}

// parses the target list from the select statement i.e all of the things to select
func (p *Parser) parseTargetList() []*Node {
	var nodes []*Node

	for p.curTok.Type != EOF && p.curTok.Type != KEYWORD && p.curTok.Literal != "FROM" {
		// skip commas in the target list
		if p.curTok.Type != COMMA {
			node := &Node{NodeType: p.curTok.Literal}
			nodes = append(nodes, node)
		}
		p.nextToken()
	}

	return nodes
}

// parses the from clause from the select
func (p *Parser) parseFromClause() []*Node {
	var nodes []*Node

	// Loop until we hit a keyword that indicates the end of the FROM clause
	for p.curTok.Type != EOF && p.curTok.Type != KEYWORD && p.curTok.Literal != "WHERE" && p.curTok.Literal != "GROUP" && p.curTok.Literal != "ORDER" && p.curTok.Literal != "LIMIT" {
		if p.curTok.Type != PERIOD {
			node := &Node{NodeType: p.curTok.Literal}
			nodes = append(nodes, node)

			// Skip commas (in case of multiple tables)
			if p.curTok.Type == COMMA {
				p.nextToken()
			}
		}
		p.nextToken()
	}

	return nodes
}

// parses the where clause from the select
func (p *Parser) parseWhereClause() *Node {
	node := &Node{NodeType: p.curTok.Literal}
	p.nextToken()
	return node
}

// peeks at the next token to check it's type, returns true if it's a TokenType
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTok.Type == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}
