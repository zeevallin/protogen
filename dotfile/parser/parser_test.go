package parser_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/parser"
)

func TestParseConfigurationFile(t *testing.T) {
	t.Run("source and language statements", func(t *testing.T) {
		input := `
		source git@github.com:zeeraw/protogen.git
		language go

		generate bar v1.0.0
		generate bar/baz v1.0.0
		generate fizz/buzz master
		generate furry/trash/can
		`
		l := lexer.New(log.New(ioutil.Discard, "", 0), []byte(input))
		p := parser.New(log.New(ioutil.Discard, "", 0), l)
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
		}
	})
}
