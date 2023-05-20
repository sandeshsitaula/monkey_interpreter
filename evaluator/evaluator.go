package evaluator

import (
	"github.com/sandeshsitaula/monkeyinter/ast"
	"github.com/sandeshsitaula/monkeyinter/object"
)

var (
	NULL  = &object.NULL{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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

		//slow impl since true and false are only two values no need to create new instance every time
		//	return &object.Boolean{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	}
	return nil
}
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmt []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmt {
		result = Eval(statement)
	}
	return result
}
