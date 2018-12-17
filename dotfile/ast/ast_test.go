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

	testGoPlugin = "grpc"
)

const expectedMinimallyViable = `source github.com/zeeraw/protogen-protos
language go {
	plugin grpc
}
output ./vendor/protos`

func TestConfigurationFile(t *testing.T) {
	t.Run("minimally viable configuration file", func(t *testing.T) {
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
					Block: &ast.Block{
						Depth: 1,
						Statements: []ast.Statement{
							&ast.GoPluginStatement{
								Token: token.Token{
									Literal: token.KWGoPlugin,
									Type:    token.GOPLUGIN,
								},
								Name: &ast.Identifier{
									Token: token.Token{
										Literal: testGoPlugin,
										Type:    token.IDENTIFIER,
									},
									Value: testGoPlugin,
								},
							},
						},
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
		test.AssertEqual(t, expectedMinimallyViable, actual.String())
	})
}
