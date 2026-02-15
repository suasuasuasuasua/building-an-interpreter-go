// Package lexer provides definition for a lexer
package lexer

import (
	"github.com/suasuasuasuasua/building-an-interpreter-go/token"
)

type Lexer struct {
	input        string // the input program
	position     int    // the current position in the input
	readPosition int    // the current reading position in the input (after ch)
	ch           byte   // the current char being examined
}

// create a new lexer with some input program
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// read the current character from the input
func (l *Lexer) readChar() {
	// reset the character to NUL if we've read more than the input
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// otherwise, update the character on the current position
		l.ch = l.input[l.readPosition]
	}

	// the position index will lag behind the read position index
	l.position = l.readPosition
	l.readPosition += 1
}

// process and return the current token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// whitespace in this programming language doesn't have any meaning. just
	// ignore it
	l.skipWhitespace()

	switch l.ch {
	case '=':
		// if the next character is an '=', then construct an EQ token
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // remember to advance the lexer
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			// else it must just be an assignment
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// if the next character is an '=', then construct an NOT_EQ token
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // remember to advance the lexer
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			// else it must just be an negation
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// first, attempt to read the identifier
			tok.Literal = l.readIdentifier()
			// second, determine the token type if it is in the list of keywords
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// find and build an identifier by continously reading characters
// this function will advance the lexer's state until it reaches a
// non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Check if a character is a valid letter in an identifier
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Check if a character is a number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// find a build a number by reading characters
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// read the input and skip any whitespace characters like spaces, tabs,
// newlines, and carriage returns.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// helper function to look at the next character while considering string length
// boundaries. note that this functinon not advance the (read)position of the
// lexer
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// helper function to create a Token given some type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
