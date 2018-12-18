package evaluator_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/evaluator"
	"github.com/zeeraw/protogen/dotfile/token"
	"github.com/zeeraw/protogen/test"
)

func TestEval(t *testing.T) {
	var logger = log.New(ioutil.Discard, "", 0)
	t.Run("when everything is alright", func(tt *testing.T) {
		tree := new()
		e := evaluator.New(logger)
		conf, err := e.Eval(tree)
		test.AssertEqual(t, nil, err)
		test.AssertEqual(t, 2, len(conf.Packages))
	})
	t.Run("when the wrong statements are made in a Go language block", func(tt *testing.T) {
		tree := new()
		tree.Statements[1] = &ast.LanguageStatement{
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
			Block: &ast.Block{
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
				},
			},
		}

		e := evaluator.New(logger)
		_, err := e.Eval(tree)
		if err == nil {
			t.Errorf("expected there to be an error")
		}
	})
}

func new() *ast.ConfigurationFile {
	return &ast.ConfigurationFile{
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
}
