package parser_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/zeeraw/protogen/test"

	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/parser"
)

func newParser(input string) *parser.Parser {
	logger := log.New(ioutil.Discard, "", 0)
	l := lexer.New(logger, []byte(input))
	return parser.New(logger, l)
}

func TestParseConfigurationFile(t *testing.T) {
	t.Run("source and language statements", func(t *testing.T) {
		p := newParser(`
			source github.com/zeeraw/protogen-protos
			language go
			generate bar v1.0.0
			generate bar/baz v1.0.0
			generate fizz/buzz master
			generate furry/trash/can
		`)
		cf, err := p.ParseConfigurationFile()
		if err != nil {
			panic(err)
		}
		if cf == nil {
			t.Fatalf("cannot be nil")
		}
		if len(cf.Statements) != 6 {
			t.Fatalf("ast.ConfigurationFile.SourceStatements does not contain 6 statement(s). got=%d", len(cf.Statements))
		}
		tests := []struct {
			t string
		}{
			{"source"},
			{"language"},
			{"generate"},
			{"generate"},
			{"generate"},
			{"generate"},
		}
		for i, tt := range tests {
			stmt := cf.Statements[i]
			lit := stmt.TokenLiteral()
			if lit != tt.t {
				t.Fatalf("token literal %s was not %s", lit, tt.t)
			}
			switch node := stmt.(type) {
			case *ast.LanguageStatement:
				if node.Block != nil {
					t.Fatalf("language statement node block be nil, was: %v", node.Block)
				}
			}
		}
	})

	t.Run("source statement and language statement with a block", func(t *testing.T) {
		p := newParser(`
			source github.com/zeeraw/protogen-protos
			language go {
				plugin grpc
				path source_relative
			}
			generate bar v1.0.0
			generate bar/baz v1.0.0
			generate fizz/buzz master
			generate furry/trash/can
		`)
		cf, err := p.ParseConfigurationFile()
		if err != nil {
			panic(err)
		}
		if cf == nil {
			t.Fatalf("cannot be nil")
		}
		if len(cf.Statements) != 6 {
			t.Fatalf("ast.ConfigurationFile.SourceStatements does not contain 6 statement(s). got=%d", len(cf.Statements))
		}
		tests := []struct {
			t string
		}{
			{"source"},
			{"language"},
			{"generate"},
			{"generate"},
			{"generate"},
			{"generate"},
		}
		for i, tt := range tests {
			stmt := cf.Statements[i]
			lit := stmt.TokenLiteral()
			if lit != tt.t {
				t.Fatalf("token literal %s was not %s", lit, tt.t)
			}
			switch node := stmt.(type) {
			case *ast.LanguageStatement:
				if node.Block == nil {
					t.Fatalf("language statement node block should not be nil")
					return
				}
				bl := len(node.Block.Statements)
				el := 2
				if bl != el {
					t.Fatalf("language block statements should have a length of %d, was: %d", el, bl)
				}
				plugin, ok := node.Block.Statements[0].(*ast.GoPluginStatement)
				if !ok {
					t.Fatalf("first language block statement should be an *ast.GoPluginStatement")
				}
				test.AssertEqual(t, "grpc", plugin.Name.String())
				path, ok := node.Block.Statements[1].(*ast.GoPathStatement)
				if !ok {
					t.Fatalf("second language block statement should be an *ast.GoPathStatement")
				}
				test.AssertEqual(t, "source_relative", path.Type.String())
			}
		}
	})
}
