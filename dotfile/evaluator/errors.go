package evaluator

import (
	"fmt"

	"github.com/zeeraw/protogen/config"
)

// ErrLanguageNotSupported happens when evaluating an ast with an invalid language
type ErrLanguageNotSupported struct {
	lang config.Language
}

func (e ErrLanguageNotSupported) Error() string {
	return fmt.Sprintf("language not supported: %q", e.lang)
}
