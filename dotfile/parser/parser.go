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
		errors: []error{},
		logger: logger,
	}
	p.Next() // Roll forward once to set the peek token
	p.Next() // Roll forward twice to make the first peek token the current token
	return p
}

// Parser represens the current parse job
type Parser struct {
	l      *lexer.Lexer
	curr   token.Token
	peek   token.Token
	errors []error

	logger *log.Logger
}

// Parse will attempt to parse a configuration file
func (p *Parser) Parse() (*ast.ConfigurationFile, error) {
	f := ast.NewConfigurationFile()
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			f.Statements = append(f.Statements, stmt)
		}
		p.Next()
	}
	var err error
	if len(p.errors) > 0 {
		err = fmt.Errorf("parsing error: there were %d errors while parsing configuration file", len(p.errors))
	}
	return f, err
}

// Error will add an error to the parser errors
func (p *Parser) Error(err error) {
	p.errors = append(p.errors, err)
}

// Errors return all current errors for the parser
func (p *Parser) Errors() []error {
	return p.errors
}

// Next will roll the parser forward by one token
func (p *Parser) Next() {
	p.curr = p.peek
	p.peek = p.l.NextToken()
}

// Peek will look at the next token without rolling the parser forward
func (p *Parser) Peek() token.Token {
	return p.peek
}

// Token will look at the current token
func (p *Parser) Token() token.Token {
	return p.curr
}

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

func (p *Parser) skipWhitespaceUntilAny(ts ...token.Type) bool {
	if !p.expectPeek(token.WHITESPACE) {
		return false
	}
	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}
	if !p.expectAny(ts...) {
		return false
	}
	return true
}

func (p *Parser) skipWhitespaceUntil(t token.Type) bool {
	if !p.expectPeek(token.WHITESPACE) {
		return false
	}
	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}
	if !p.expect(t) {
		return false
	}
	return true
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.Token().Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.Peek().Type == t
}

func (p *Parser) expectAny(ts ...token.Type) bool {
	for _, t := range ts {
		if p.curTokenIs(t) {
			return true
		}
	}
	p.Next()
	return false
}

func (p *Parser) expect(t token.Type) bool {
	if p.curTokenIs(t) {
		return true
	}
	p.Next()
	return false
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
	msg := fmt.Errorf("expected next token to be %s, got %s instead", t, p.Peek().Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) isTerminus() bool {
	return p.curTokenIs(token.NEWLINE) || p.curTokenIs(token.EOF)
}

func (p *Parser) isBlockStart() bool {
	return p.curTokenIs(token.LEFTBRACE)
}
