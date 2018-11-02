package parser

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/token"
)

// New prepares and returns a parser for a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

// Parser represens the current parse job
type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

// ParseConfigurationFile will attempt to parse a configuration file
func (p *Parser) ParseConfigurationFile() (cf *ast.ConfigurationFile, err error) {
	cf = &ast.ConfigurationFile{}
	cf.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			cf.Statements = append(cf.Statements, stmt)
		}
		p.nextToken()
	}

	spew.Dump(cf)

	return cf, nil
}

// Errors return all current errors for the parser
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.SOURCE:
		return p.parseSourceStatement()
	case token.LANGUAGE:
		return p.parseLanguageStatement()
	case token.GENERATE:
		return p.parseGenerateStatement()
	default:
		return nil
	}
}

func (p *Parser) parseSourceStatement() ast.Statement {
	stmt := &ast.SourceStatement{
		Token: p.curToken,
	}
	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Source = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseLanguageStatement() ast.Statement {
	stmt := &ast.LanguageStatement{
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

	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseGenerateStatement() ast.Statement {
	stmt := &ast.GenerateStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Target = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

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

	idnt := ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if p.curTokenIs(token.VERSION) {
		stmt.Tag = &ast.Version{
			Identifier: idnt,
		}
	} else {
		stmt.Tag = &idnt
	}

	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectAnyPeek(ts ...token.Type) bool {
	for _, t := range ts {
		if p.peekTokenIs(t) {
			p.nextToken()
			return true
		}
	}
	return false
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
