package parser

import (
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.Token().Type {
	case token.SOURCE:
		return p.parseSourceStatement()
	case token.LANGUAGE:
		return p.parseLanguageStatement()
	case token.OUTPUT:
		return p.parseOutputStatement()
	case token.GENERATE:
		return p.parseGenerateStatement()
	case token.PATH:
		return p.parsePathStatement()
	case token.PLUGIN:
		return p.parsePluginStatement()
	default:
		return nil
	}
}

func (p *Parser) parseBlock() *ast.Block {
	var statements []ast.Statement
	for !p.curTokenIs(token.RIGHTBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.Next()
	}
	return &ast.Block{
		Statements: statements,
	}
}

func (p *Parser) parseSourceStatement() ast.Statement {
	stmt := &ast.SourceStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Source = ast.NewIdentifier(p.Token())
	for !p.isTerminus() {
		p.Next()
	}
	return stmt
}

func (p *Parser) parseLanguageStatement() ast.Statement {
	stmt := &ast.LanguageStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ast.NewIdentifier(p.Token())
	p.Next()
	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}
	if p.isBlockStart() {
		stmt.Block = p.parseBlock()
	}
	for !p.isTerminus() {
		p.Next()
	}
	return stmt
}

func (p *Parser) parseOutputStatement() ast.Statement {
	stmt := &ast.OutputStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Path = ast.NewIdentifier(p.Token())
	for !p.isTerminus() {
		p.Next()
	}
	return stmt
}

func (p *Parser) parseGenerateStatement() ast.Statement {
	stmt := &ast.GenerateStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Target = ast.NewIdentifier(p.Token())
	if p.peekTokenIs(token.NEWLINE) {
		stmt.Tag = nil
		return stmt
	}
	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectAnyPeek(token.IDENTIFIER, token.VERSION) {
		return nil
	}
	idnt := ast.NewIdentifier(p.Token())
	if p.curTokenIs(token.VERSION) {
		stmt.Tag = &ast.Version{
			Identifier: *idnt,
		}
	} else {
		stmt.Tag = idnt
	}
	for !p.isTerminus() {
		p.Next()
	}
	return stmt
}

func (p *Parser) parsePathStatement() ast.Statement {
	stmt := &ast.PathStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Type = ast.NewIdentifier(p.Token())
	return stmt
}

func (p *Parser) parsePluginStatement() ast.Statement {
	stmt := &ast.PluginStatement{
		Token: p.Token(),
	}
	if !p.skipWhitespaceUntil(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ast.NewIdentifier(p.Token())
	return stmt
}
