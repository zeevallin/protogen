package evaluator

import (
	"fmt"

	"github.com/zeeraw/protogen/dotfile/ast"

	"github.com/zeeraw/protogen/config"
)

// ErrLanguageNotSupported happens when evaluating an ast with an invalid language
type ErrLanguageNotSupported struct {
	lang config.Language
}

func (e ErrLanguageNotSupported) Error() string {
	return fmt.Sprintf("language not supported: %q", e.lang)
}

// ErrStatementNotSupported happens when a statement is not supported for a given language
type ErrStatementNotSupported struct {
	lang      config.Language
	statement ast.Statement
}

func (e ErrStatementNotSupported) Error() string {
	return fmt.Sprintf("statement not supported for %q language: %T", e.lang, e.statement)
}
