package ast

import (
	"bytes"
	"phour/token"
	"strings"
)

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

// IntegerLiteral ..
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral ..
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression ..
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral ..
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression ..
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (pe *InfixExpression) expressionNode() {}

// TokenLiteral ..
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(" " + pe.Operator + " ")
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean .. Type declaration to handle Booleans
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral ..
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IfExpression .. Type declaration to handle If expressions
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral ..
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// BlockStatement ..
type BlockStatement struct {
	Token      token.Token // { Left curly token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral ..
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral to support parsing of Functions
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f1 *FunctionLiteral) expressionNode() {}

// TokenLiteral ..
func (f1 *FunctionLiteral) TokenLiteral() string { return f1.Token.Literal }
func (f1 *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f1.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f1.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f1.Body.String())

	return out.String()
}

// CallExpression ..
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral ..
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
