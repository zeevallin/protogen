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

	testGoPlugin    = "grpc"
	testOptionName  = "WithKey"
	testOptionValue = "AndValue"
)

const expectedMinimallyViable = `source github.com/zeeraw/protogen-protos
language go {
	plugin grpc
	option WithKey AndValue
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
							&ast.PluginStatement{
								Token: token.Token{
									Literal: token.KWPlugin,
									Type:    token.PLUGIN,
								},
								Name: &ast.Identifier{
									Token: token.Token{
										Literal: testGoPlugin,
										Type:    token.IDENTIFIER,
									},
									Value: testGoPlugin,
								},
							},
							&ast.OptionStatement{
								Token: token.Token{
									Literal: token.KWOption,
									Type:    token.OPTION,
								},
								Name: &ast.Identifier{
									Token: token.Token{
										Literal: testOptionName,
										Type:    token.IDENTIFIER,
									},
									Value: testOptionName,
								},
								Value: &ast.Identifier{
									Token: token.Token{
										Literal: testOptionValue,
										Type:    token.IDENTIFIER,
									},
									Value: testOptionValue,
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
