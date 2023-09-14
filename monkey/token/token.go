package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 12343

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyworkds
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Using a string for ease of use and debugging.
// Not as porformant but perfectly fine
// Could be an int or byte if performance was an issue
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
