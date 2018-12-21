package lexer_test

import (
	"testing"

	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/token"
)

func TestNextToken(t *testing.T) {
	input := "source github.com/foo/bar-lol-wat language go output bar\nv0.0.1\nsource sourcelanguage js;[]{}()"
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.SOURCE, "source"},
		{token.WHITESPACE, " "},
		{token.IDENTIFIER, "github.com/foo/bar-lol-wat"},
		{token.WHITESPACE, " "},
		{token.LANGUAGE, "language"},
		{token.WHITESPACE, " "},
		{token.IDENTIFIER, "go"},
		{token.WHITESPACE, " "},
		{token.OUTPUT, "output"},
		{token.WHITESPACE, " "},
		{token.IDENTIFIER, "bar"},
		{token.NEWLINE, "\n"},
		{token.VERSION, "v0.0.1"},
		{token.NEWLINE, "\n"},
		{token.SOURCE, "source"},
		{token.WHITESPACE, " "},
		{token.IDENTIFIER, "sourcelanguage"},
		{token.WHITESPACE, " "},
		{token.IDENTIFIER, "js"},
		{token.ILLEGAL, ";"},
		{token.ILLEGAL, "["},
		{token.ILLEGAL, "]"},
		{token.LEFTBRACE, "{"},
		{token.RIGHTBRACE, "}"},
		{token.ILLEGAL, "("},
		{token.ILLEGAL, ")"},
		{token.EOF, ""},
	}
	l := lexer.New([]byte(input))

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
