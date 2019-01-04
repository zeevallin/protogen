package parser

import "github.com/zeeraw/protogen/dotfile/token"

func (p *Parser) skipWhitespaceUntilAny(ts ...token.Type) bool {
	if !p.expectPeek(token.WHITESPACE) {
		return false
	}
	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}
	return p.expectAny(ts...)
}

func (p *Parser) skipWhitespaceUntil(t token.Type) bool {
	if !p.expectPeek(token.WHITESPACE) {
		return false
	}
	for p.curTokenIs(token.WHITESPACE) {
		p.Next()
	}
	return p.expect(t)
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

func (p *Parser) isTerminus() bool {
	return p.curTokenIs(token.NEWLINE) || p.curTokenIs(token.EOF)
}

func (p *Parser) isBlockStart() bool {
	return p.curTokenIs(token.LEFTBRACE)
}

func (p *Parser) isBlockEnd() bool {
	return p.curTokenIs(token.RIGHTBRACE)
}
