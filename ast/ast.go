package ast

import "github.com/sandeshsitaula/monkeyinter/token"

// Representing all statements as Node
type Node interface {
	TokenLiteral() string
}

// Represents Statement like Let Return
type Statement interface {
	Node
	statementNode()
}

// Represents Expression like 10*5,10+5
type Expression interface {
	Node
	expressionNode()
}

// ast is the colelction of statements
type Program struct {
	Statements []Statement
}

// Get the literal for Statement For LET Statement literal is "let"
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// for Return
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) expressionNode()      {}
func (l *Identifier) TokenLiteral() string { return l.Token.Literal }
