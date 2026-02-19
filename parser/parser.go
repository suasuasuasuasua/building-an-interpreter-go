// Package parser provides an implementation of a parser
package parser

import (
	"fmt"

	"github.com/suasuasuasuasua/building-an-interpreter-go/ast"
	"github.com/suasuasuasuasua/building-an-interpreter-go/lexer"
	"github.com/suasuasuasuasua/building-an-interpreter-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

// create a new parser given a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// peek two tokens so we can fill the current and peeked token
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s. got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// update the current token and peek the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() // check the next token in the lexer
}

// parse the lines in a program and assemble a list of statements
// run until the EOF token marker
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// at this step, the current token == token.LET
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// at this step, the current token == token.IDENT
	// and the value is the literal name of the variable
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// expect an assignment operator
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: placeholder for handling values in the future
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// check if the current token is some type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// check if the next token is some type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// check what the next token is. if it is type `t`, then advance the state of
// the parser and  lexer and return true
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
