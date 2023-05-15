package parser

import (
	"testing"

	"github.com/sandeshsitaula/monkeyinter/ast"
	"github.com/sandeshsitaula/monkeyinter/lexer"
)

func TestLetStatements(t *testing.T) {
	/*	input := `
		let x=5;
		let y=10;
		let foobar=84848;
			`*/
	input := `
	let x 5;
	let =10;
	let 8838383;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program statements doesnot contain 3 statements got %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("not 'let' but got =%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement,got %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letstmt.name.value not %s ,got %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error : %q", msg)
	}
	t.FailNow()
}
