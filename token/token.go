// Package token provides definitions for tokens in the monkey language
package token

// alias tokens as strings
//
// NOTE: this is somewhat of a naiive implementation since tokens can be
// represented with smaller types like bytes or integers. for the purpose of
// this book, a string gives us flexibility and debuggability
type TokenType string

// a token has some "type" and the associated literal
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
