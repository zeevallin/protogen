package parser

import (
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/token"
)

func (p *Parser) parseGoStatement() ast.Statement {
	switch p.curToken.Type {
	case token.PATH:
		return p.parseGoPathStatement()
	case token.PLUGIN:
		return p.parseGoPluginStatement()
	default:
		return nil
	}
}

func (p *Parser) parseGoPathStatement() ast.Statement {
	stmt := &ast.GoPathStatement{
		Token: p.curToken,
	}
	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Type = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	return stmt
}

func (p *Parser) parseGoPluginStatement() ast.Statement {
	stmt := &ast.GoPluginStatement{
		Token: p.curToken,
	}
	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	return stmt
}
