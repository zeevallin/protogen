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
