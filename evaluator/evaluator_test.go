package evaluator

import (
	"testing"

	"github.com/suasuasuasuasua/building-an-interpreter-go/ast"
	"github.com/suasuasuasuasua/building-an-interpreter-go/lexer"
	"github.com/suasuasuasuasua/building-an-interpreter-go/object"
	"github.com/suasuasuasuasua/building-an-interpreter-go/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range testCases {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not an Integer. get=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has the wrong value. got=%d, want=%d", result.Value, expected)
	}

	return true
}

func TestEval(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		node ast.Node
		want object.Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Eval(tt.node)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
