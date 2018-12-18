package parser

import (
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/token"
)

func (p *Parser) parseGoStatement() ast.Statement {
	switch p.Token().Type {
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
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Type = ast.NewIdentifier(p.Token())
	return stmt
}

func (p *Parser) parseGoPluginStatement() ast.Statement {
	stmt := &ast.GoPluginStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ast.NewIdentifier(p.Token())
	return stmt
}
