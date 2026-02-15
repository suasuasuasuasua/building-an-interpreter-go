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

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// helper function to create a Token given some type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
