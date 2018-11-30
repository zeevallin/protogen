package token_test

import (
	"testing"

	"github.com/zeeraw/protogen/dotfile/token"
	"github.com/zeeraw/protogen/test"
)

func TestLookupIdentifier(t *testing.T) {
	t.Run("with known identifiers", func(tt *testing.T) {
		test.AssertEqual(t, token.Type("SOURCE"), token.LookupIdentifier("source"))
		test.AssertEqual(t, token.Type("LANGUAGE"), token.LookupIdentifier("language"))
		test.AssertEqual(t, token.Type("GENERATE"), token.LookupIdentifier("generate"))
		test.AssertEqual(t, token.Type("OUTPUT"), token.LookupIdentifier("output"))
	})
	t.Run("with unknown identifier", func(tt *testing.T) {
		test.AssertEqual(t, token.Type("IDENTIFIER"), token.LookupIdentifier("some/thing"))
		test.AssertEqual(t, token.Type("IDENTIFIER"), token.LookupIdentifier("foobar"))
		test.AssertEqual(t, token.Type("IDENTIFIER"), token.LookupIdentifier("hello!"))
	})
}
