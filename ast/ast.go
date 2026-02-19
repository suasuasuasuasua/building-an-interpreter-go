// Package ast provides definitions for an abstract syntax tree
package ast

import (
	"github.com/suasuasuasuasua/building-an-interpreter-go/token"
)

// the literal value that the node is associated with
type Node interface {
	TokenLiteral() string
}

// a statement node
type Statement interface {
	Node
	statementNode()
}

// an expression node
type Expression interface {
	Node
	expressionNode()
}

// a program is a series of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.Let token
	Name  *Identifier // the identifier in the statement
	Value Expression  // the expression given to the identifier
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the name of the identifier
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
