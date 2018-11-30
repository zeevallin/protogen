package ast_test

import (
	"testing"

	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/dotfile/token"
	"github.com/zeeraw/protogen/test"
)

const (
	testSource = "github.com/zeeraw/protogen-protos"
	testOutput = "./vendor/protos"
	testLang   = "go"
)

func TestConfigurationFile(t *testing.T) {
	t.Run("minimally viable configuration file", func(t *testing.T) {
		expected := "source github.com/zeeraw/protogen-protos\nlanguage go\noutput ./vendor/protos"
		actual := &ast.ConfigurationFile{
			Statements: []ast.Statement{
				&ast.SourceStatement{
					Token: token.Token{
						Literal: token.KWSource,
						Type:    token.SOURCE,
					},
					Source: &ast.Identifier{
						Token: token.Token{
							Literal: testSource,
							Type:    token.IDENTIFIER,
						},
						Value: testSource,
					},
				},
				&ast.LanguageStatement{
					Token: token.Token{
						Literal: token.KWLanguage,
						Type:    token.LANGUAGE,
					},
					Name: &ast.Identifier{
						Token: token.Token{
							Literal: testLang,
							Type:    token.IDENTIFIER,
						},
						Value: testLang,
					},
				},
				&ast.OutputStatement{
					Token: token.Token{
						Literal: token.KWOutput,
						Type:    token.OUTPUT,
					},
					Path: &ast.Identifier{
						Token: token.Token{
							Literal: testOutput,
							Type:    token.IDENTIFIER,
						},
						Value: testOutput,
					},
				},
			},
		}
		test.AssertEqual(t, expected, actual.String())
	})
}
