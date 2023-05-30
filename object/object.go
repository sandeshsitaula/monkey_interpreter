// object package represents typesystem like Integer,Boolean
package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sandeshsitaula/monkeyinter/ast"
)
//for creating builtin functions for our interpreter
type BuiltinFunction func(args ...Object)Object


type ObjectType string

const (

	BUILTIN_OBJ="BUILTIN"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	ERROR_OBJ        = "ERROR"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	FUNCTION_OBJ     = "FUNCTION"
	//for string  evaluation
	STRING_OBJ="STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// represents integer and satisfies Object interface
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type NULL struct{}

func (n *NULL) Type() ObjectType { return NULL_OBJ }
func (n *NULL) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// to keep track of value bindings
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}


//for storing string data
type String struct{
	Value string
}

func (s *String)Type()ObjectType{return STRING_OBJ}
func (s *String )Inspect()string{return s.Value}

type Builtin struct{
	Fn BuiltinFunction
}


func (b *Builtin)Type()ObjectType{return BUILTIN_OBJ}
func (b *Builtin)Inspect()string{return "builtin function"}
