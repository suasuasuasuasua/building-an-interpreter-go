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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "=="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// define the unique keywords in the language
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// helper function that determines if an identifier is a keyword or just an
// identifier
func LookupIdent(ident string) TokenType {
	// in go, performing a map lookup returns two values, the value and an error
	// flag. we can check the error flag in the same line
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	// if we can't find the identifier in the map, then it must be an identifier
	return IDENT
}
