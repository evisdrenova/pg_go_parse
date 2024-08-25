package pggoquery

import (
	"unicode"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	currentChar  byte // current char being examined
}

// returns a new instance of a lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// reads the current character and assigns it to the ch field in the Lexer struct and increments to the next character
// Lexer works with two pointers to characters, where the position is the current char and readPosition is the next character that's being read
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// reads in the next token and returns a new token based on the char
func (l *Lexer) NextToken() Token {

	var tok Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		tok = newToken(ASSIGN, l.currentChar)
	case '+':
		tok = newToken(PLUS, l.currentChar)
	case '-':
		tok = newToken(MINUS, l.currentChar)
	case '!':
		tok = newToken(BANG, l.currentChar)
	case '*':
		tok = newToken(ASTERISK, l.currentChar)
	case '/':
		tok = newToken(SLASH, l.currentChar)
	case '<':
		tok = newToken(LT, l.currentChar)
	case '>':
		tok = newToken(GT, l.currentChar)
	case ',':
		tok = newToken(COMMA, l.currentChar)
	case ';':
		tok = newToken(SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(LPAREN, l.currentChar)
	case ')':
		tok = newToken(RPAREN, l.currentChar)
	case '{':
		tok = newToken(LBRACE, l.currentChar)
	case '}':
		tok = newToken(RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = IDENT
			if isKeyword(tok.Literal) {
				tok.Type = KEYWORD
			}
			return tok
		} else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = INT
			return tok
		} else {
			tok = newToken(ILLEGAL, l.currentChar)
		}
	}
	l.readChar()
	return tok
}

// returns the identifier from the input
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// reads the number from the input
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// skips a whitespace if it finds one
func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

// returns a new token instance based on the input
func newToken(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}

// checks if the char is a letter
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// checks if the char is a digit
func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

// checks if the identifier is a keyword
func isKeyword(ident string) bool {
	keywords := []string{"SELECT", "FROM", "WHERE", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "TABLE", "INTO"} // we can probably add more here
	for _, keyword := range keywords {
		if ident == keyword {
			return true
		}
	}
	return false
}
