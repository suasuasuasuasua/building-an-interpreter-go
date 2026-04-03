package evaluator

import (
	"fmt"

	"github.com/suasuasuasuasua/building-an-interpreter-go/ast"
	"github.com/suasuasuasuasua/building-an-interpreter-go/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BlockStatement:
	case *ast.Boolean:
	case *ast.CallExpression:
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.FunctionLiteral:
	case *ast.Identifier:
	case *ast.IfExpression:
	case *ast.InfixExpression:
	case *ast.LetStatement:
	case *ast.PrefixExpression:
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ReturnStatement:
	default:
		panic(fmt.Sprintf("unexpected ast.Node: %#v", node))
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
