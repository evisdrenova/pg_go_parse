package pggoquery

type TokenType string

// defines the sql tokens that we might encounter in a sql query
const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF                 = "EOF"
	IDENT               = "IDENT"
	INT                 = "INT"
	FLOAT               = "FLOAT"
	ASSIGN              = "="
	PLUS                = "+"
	MINUS               = "-"
	BANG                = "!"
	ASTERISK            = "*"
	SLASH               = "/"
	LT                  = "<"
	GT                  = ">"
	COMMA               = ","
	SEMICOLON           = ";"
	LPAREN              = "("
	RPAREN              = ")"
	LBRACE              = "{"
	RBRACE              = "}"
	KEYWORD             = "KEYWORD"
)

type Token struct {
	Type    TokenType
	Literal string
}
