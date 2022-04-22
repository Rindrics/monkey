package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	ASSIGN = "ASSIGN"
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	LPAREN = ")"
	LBRACE = "{"
	LBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
