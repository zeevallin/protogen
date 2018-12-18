package parser

import (
	"fmt"
	"log"

	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/token"
)

// New prepares and returns a parser for a lexer
func New(logger *log.Logger, l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
		logger: logger,
	}
	p.Next() // Roll forward once to set the peek token
	p.Next() // Roll forward twice to make the first peek token the current token
	return p
}

// Parser represens the current parse job
type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
	logger    *log.Logger
}

// Parse will attempt to parse a configuration file
func (p *Parser) Parse() (cf *ast.ConfigurationFile, err error) {
	cf = &ast.ConfigurationFile{}
	cf.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			cf.Statements = append(cf.Statements, stmt)
		}
		p.Next()
	}

	return cf, nil
}

// Errors return all current errors for the parser
func (p *Parser) Errors() []string {
	return p.errors
}

// Next will roll the parser forward by one token
func (p *Parser) Next() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.SOURCE:
		return p.parseSourceStatement()
	case token.LANGUAGE:
		return p.parseLanguageStatement()
	case token.OUTPUT:
		return p.parseOutputStatement()
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

	for !p.isTerminus() {
		p.Next()
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
	p.Next()

	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}

	if p.isBlockStart() {
		stmt.Block = p.parseBlock()
	}

	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.Next()
	}

	return stmt
}

func (p *Parser) parseBlock() *ast.Block {
	var statements []ast.Statement
	for p.curToken.Type != token.RIGHTBRACE {
		stmt := p.parseGoStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.Next()
	}
	return &ast.Block{
		Statements: statements,
	}
}

func (p *Parser) parseOutputStatement() ast.Statement {
	stmt := &ast.OutputStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.WHITESPACE) {
		return nil
	}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Path = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.Next()
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
		p.Next()
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
			p.Next()
			return true
		}
	}
	return false
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.Next()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) isBlockStart() bool {
	return p.curTokenIs(token.LEFTBRACE)
}
