package parser

import (
	"fmt"

	"github.com/zeeraw/protogen/dotfile/token"
)

// NewParsingErr returns a parsing error
func NewParsingErr(line, col int, msg string) error {
	return ErrParsing{line, col, msg}
}

// ErrParsing is an error that happens when
type ErrParsing struct {
	line int
	col  int
	msg  string
}

func (e ErrParsing) Error() string {
	return fmt.Sprintf("parsing error line=%d col=%d: %s", e.line, e.col, e.msg)
}

func (p *Parser) peekError(t token.Type) {
	err := fmt.Errorf("expected next token to be %s, got %s instead", t, p.Peek().Type)
	p.Error(err)
}
