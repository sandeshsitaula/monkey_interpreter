package evaluator

import (
	"github.com/sandeshsitaula/monkeyinter/ast"
	"github.com/sandeshsitaula/monkeyinter/object"
)

// Eval takes ast.Node generted by our parser and convers to object Type
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

		//Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmt []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmt {
		result = Eval(statement)
	}
	return result
}
