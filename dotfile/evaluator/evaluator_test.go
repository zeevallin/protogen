package evaluator_test

import (
	"testing"

	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/evaluator"
	"github.com/zeeraw/protogen/dotfile/token"
	"github.com/zeeraw/protogen/test"
)

func TestEval(t *testing.T) {
	t.Run("a configuration file with two generation statements", func(tt *testing.T) {
		tree := &ast.ConfigurationFile{
			Statements: []ast.Statement{
				&ast.SourceStatement{
					Token: token.Token{
						Type:    token.SOURCE,
						Literal: "source",
					},
					Source: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENTIFIER,
							Literal: "github.com/zeeraw/protogen-protos",
						},
						Value: "github.com/zeeraw/protogen-protos",
					},
				},
				&ast.LanguageStatement{
					Token: token.Token{
						Type:    token.LANGUAGE,
						Literal: "language",
					},
					Name: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENTIFIER,
							Literal: "go",
						},
						Value: "go",
					},
				},
				&ast.GenerateStatement{
					Token: token.Token{
						Type:    token.GENERATE,
						Literal: "generate",
					},
					Target: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENTIFIER,
							Literal: "foo/bar",
						},
						Value: "foo/bar",
					},
					Tag: &ast.Identifier{},
				},
				&ast.GenerateStatement{
					Token: token.Token{
						Type:    token.GENERATE,
						Literal: "generate",
					},
					Target: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENTIFIER,
							Literal: "fizz/buzz",
						},
						Value: "fizz/buzz",
					},
					Tag: &ast.Version{},
				},
			},
		}

		e := evaluator.New()
		conf, err := e.Eval(tree)
		test.AssertEqual(t, nil, err)
		test.AssertEqual(t, 2, len(conf.Packages))
	})
}
