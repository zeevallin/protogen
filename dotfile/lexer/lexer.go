package lexer

import (
	"log"

	"github.com/zeeraw/protogen/dotfile/token"
)

// Lexer represents a current lexing session of file
type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
	logger       *log.Logger
}

// New spawns a new lexer using the given input
func New(logger *log.Logger, input []byte) *Lexer {
	l := &Lexer{input: []rune(string(input))}
	l.readChar()
	return l
}

// NextToken will attempt to extract the next token based on the lexer's position
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '{':
		tok = newToken(token.LEFTBRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.ch)
	case ' ', '\t', '\r':
		tok = newToken(token.WHITESPACE, l.ch)
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if l.ch == 'v' && isNumber(l.peekChar()) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.VERSION
			return tok
		} else if isValidRune(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isValidRune(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isNumber(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isNewLine(ch rune) bool {
	return ch == '\n' || ch == ';'
}

func isEOF(ch rune) bool {
	return ch == '\x00'
}

func isBracket(ch rune) bool {
	return ch == '[' || ch == ']' || ch == '(' || ch == ')'
}

func isValidRune(ch rune) bool {
	return !(isWhiteSpace(ch) || isNewLine(ch) || isEOF(ch) || isBracket(ch))
}

func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
