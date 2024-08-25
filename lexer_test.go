package pggoquery

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `!=+(){},; SELECT * FROM users WHERE age > 20`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{BANG, "!"},
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{KEYWORD, "SELECT"},
		{ASTERISK, "*"},
		{KEYWORD, "FROM"},
		{IDENT, "users"},
		{KEYWORD, "WHERE"},
		{IDENT, "age"},
		{GT, ">"},
		{INT, "20"},
		{EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
