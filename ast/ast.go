package ast

import "phour/token"

import "bytes"

// Node represents each node for our AST which will get
// represented by its's token literal
type Node interface {
	TokenLiteral() string
	String() string // print AST nodes for debugging
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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
	Value Expression
}

func (r *ReturnStatement) statementNode() {}

// TokenLiteral ..
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

// ExpressionStatement struct contains Token which
// every single Node has Expression which will hold
// the expression
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral ..
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral() + " ")

	if r.Value != nil {
		out.WriteString(r.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (i *Identifier) String() string { return i.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral ..
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
