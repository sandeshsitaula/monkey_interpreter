package evaluator

import (
	"testing"

	"github.com/sandeshsitaula/monkeyinter/lexer"
	"github.com/sandeshsitaula/monkeyinter/object"
	"github.com/sandeshsitaula/monkeyinter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-9", -9},
		{"5+5+5", 15},
		{"4*5", 20},
		{"2*(2+4)", 12},
		{"50/2*2+10", 60},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
func TestEvalBoolExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"2 <5", true},
		{"5>5", false},
		{"1==1", true},
		{"2!=2", false},
		{"1!=2", true},
		{"true==true", true},
		{"true==false", false},
		{"true!=false", true},
		{"(1<4)==true", true},
		{"(1>3)==false", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not integer .got=%T(%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value.got=%d,want=%d", result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not boolean .got =%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value.got=%t,want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!23", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) {10}", 10},
		{"if (false) { 10 }", nil},
		{"if (1){10}", 10},
		{"if (1>2){10}", nil},
		{"if (1>2){10}else{20}", 20},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL.got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		//{"return 10;", 10},
		//	{"return 10;9;", 10},
		//	{"return 2*5;9;", 10},
		//	{"9;return 2*5;5;", 10},
		{`if (10>1){
			if(10>1){
			return 10;
			}
			return 1;
		}`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
if (10 > 1) {
return true + false;
}
return 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foobar", "identifier not found: foobar"},
		{
			`"hello"-"World"`,
			"unknown operator:STRING - STRING",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object for tt.input %s returned. got=%T(%+v)", tt.input,
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5;a;", 5},
		{"let a = 5*5;a;", 25},
		{"let a=5;let b=a;b;", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}


//test for string evaluation

func TestStringLiteral(t *testing.T){
input:=`"hello world"`
evaluated:=testEval(input)
str,ok:=evaluated.(*object.String)
if !ok{
	t.Fatalf("object is not string.got=%T(%+v)",evaluated,evaluated)
}
if str.Value!="hello world"{
	t.Errorf("String has wrong value.ogt=%q",str.Value)
}

}

//testing string concatenation
func TestStringConcatenation(t *testing.T){
	input:=`"Helo"+""+"World"`

	evaluated:=testEval(input)
	str,ok:=evaluated.(*object.String)

	if !ok{
		t.Fatalf("object is not String.got=%T (%+v)",evaluated,evaluated)
	}
	if str.Value!="HeloWorld"{
		t.Errorf("String has wrong value.got=%q",str.Value)
	}
}
