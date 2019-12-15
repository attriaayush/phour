package parser

import (
	"phour/ast"
	"phour/lexer"
	"phour/token"
)

// Parser ..
type Parser struct {
	l *lexer.Lexer // pointer to the instance of the lexer

	curToken  token.Token // current token under examination
	peekToken token.Token // look at the next token if curToken does not give us enough info
}

// New to get us a new token for curToken and peekToken
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

// helper function for New()
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram ..
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
