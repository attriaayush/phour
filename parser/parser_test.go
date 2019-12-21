package parser

import "testing"

import "phour/lexer"

import "phour/ast"

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let z = 50;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf(`ParseProgram() returned nil`)
	}
	if len(program.Statements) != 4 {
		t.Fatalf(`program.Statements does not contain 3 statments. go=%d`, len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf(`s.TokenLiteral not 'let' but got=%T`, s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf(`s not *ast.LetStatement. got=%T`, s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf(`letStmt.Name.Value not '%s'. But got=%s`, name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf(`letStmt.Name.TokenLiteral() not '%s'. But got=%s`, name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf(`parser has %d errors`, len(errors))
	for _, msg := range errors {
		t.Errorf(`parser error: %q`, msg)
	}
	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10; 
	return 9933322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf(`program.Statements does not contain 3 statements. got=%d`, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf(`stmt not *ast.ReturnStatement. got=%d`, stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf(`returnStmt.TokenLiteral() not 'return', got=%q`, returnStmt.TokenLiteral())
		}
	}
}
