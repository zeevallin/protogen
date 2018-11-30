package ast

import (
	"fmt"
	"strings"

	"github.com/zeeraw/protogen/dotfile/token"
)

// ConfigurationFile is the top level of our AST
type ConfigurationFile struct {
	Statements []Statement
}

// TokenLiteral returns the configuration file token literal string
func (cf *ConfigurationFile) TokenLiteral() string { return token.CONFIGURATION }
func (cf *ConfigurationFile) String() string {
	lines := make([]string, len(cf.Statements))
	for idx, stmt := range cf.Statements {
		lines[idx] = stmt.String()
	}
	return strings.Join(lines, "\n")
}

// Node is the generic interface every entity has to conform to
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a full statement in the file
type Statement interface {
	Node
}

// Expression is part of a statement
type Expression interface {
	Node
}

// SourceStatement describes the source of the proto buffers
type SourceStatement struct {
	Token  token.Token // token.SOURCE
	Source Expression
}

// TokenLiteral returns the source statement token literal string
func (ss *SourceStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SourceStatement) String() string {
	return fmt.Sprintf("%s %s", token.KWSource, ss.Source)
}

// GenerateStatement describes the proto packages to be generated
type GenerateStatement struct {
	Token  token.Token // token.GENERATE
	Target Expression
	Tag    Expression
}

// TokenLiteral returns the generate statement token literal string
func (gs *GenerateStatement) TokenLiteral() string { return gs.Token.Literal }
func (gs *GenerateStatement) String() string {
	return fmt.Sprintf("%s %s %s", token.KWGenerate, gs.Target, gs.Tag)
}

// LanguageStatement describes what language to generate the protobuffers for
type LanguageStatement struct {
	Token token.Token // token.LANGUAGE
	Name  Expression
}

// TokenLiteral returns the language statement token literal string
func (ls *LanguageStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LanguageStatement) String() string       { return fmt.Sprintf("%s %s", token.KWLanguage, ls.Name) }

// OutputStatement describes where to generate the protobuffers to
type OutputStatement struct {
	Token token.Token // token.LANGUAGE
	Path  Expression
}

// TokenLiteral returns the language statement token literal string
func (ls *OutputStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *OutputStatement) String() string       { return fmt.Sprintf("%s %s", token.KWOutput, ls.Path) }

// Identifier describes a value in the configuration
type Identifier struct {
	Token token.Token // token.IDENTIFIER
	Value string
}

// TokenLiteral returns the identifier token literal string
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Version describes a version value
type Version struct {
	Identifier
	Major int
	Minor int
	Patch int
}
