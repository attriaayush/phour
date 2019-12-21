package ast

import "phour/token"

// Node represents each node for our AST which will get
// represented by its's token literal
type Node interface {
	TokenLiteral() string
}

// Statement .. when the parser comes across any statement
type Statement interface {
	Node
	statementNode()
}

// Expression .. when the parser comes across any expression
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of our AST that parser produces
type Program struct {
	Statements []Statement
}

// TokenLiteral ..
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement ..
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral ..
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier ..
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ..
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement ..
type ReturnStatement struct {
	Token token.Token
	Value string
}

func (r *ReturnStatement) statementNode() {}

// TokenLiteral ..
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }
