package ast

import (
	"fmt"

	"github.com/zeeraw/protogen/dotfile/token"
)

// GoPluginStatement is a statement for go
type GoPluginStatement struct {
	Token token.Token // token.GOPLUGIN
	Name  Expression
}

// TokenLiteral returns the source statement token literal string
func (gps *GoPluginStatement) TokenLiteral() string { return gps.Token.Literal }
func (gps *GoPluginStatement) String() string {
	return fmt.Sprintf("%s %s", token.KWGoPlugin, gps.Name)
}

// GoPathStatement is a statement for go
type GoPathStatement struct {
	Token token.Token // token.GOPATH
	Type  Expression
}

// TokenLiteral returns the source statement token literal string
func (gps *GoPathStatement) TokenLiteral() string { return gps.Token.Literal }
func (gps *GoPathStatement) String() string {
	return fmt.Sprintf("%s %s", token.KWGoPath, gps.Type)
}
