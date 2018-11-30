package evaluator

import (
	"log"
	"strings"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/source"
)

// New returns a new evaluation instance
func New(logger *log.Logger) *Evaluator {
	return &Evaluator{
		Packages: make([]*config.Package, 0),
		logger:   logger,
	}
}

// Evaluator evaluates a config node
type Evaluator struct {
	Packages        []*config.Package
	CurrentLanguage config.Language
	CurrentSource   source.Source
	CurrentOutput   string
	logger          *log.Logger
}

// Eval evaluates a configuration file AST
func (e *Evaluator) Eval(node *ast.ConfigurationFile) (*config.Config, error) {
	err := e.eval(node)
	if err != nil {
		return nil, err
	}
	return &config.Config{
		Packages: e.Packages,
	}, nil
}

func (e *Evaluator) eval(node ast.Node) error {
	var err error
	switch node := node.(type) {
	case *ast.ConfigurationFile:
		err = e.evalStatements(node.Statements)
	case *ast.SourceStatement:
		e.CurrentSource, err = e.evalSourceStatement(node)
	case *ast.LanguageStatement:
		e.CurrentLanguage = e.evalLanguageStatement(node)
	case *ast.OutputStatement:
		e.CurrentOutput = e.evalOutputStatement(node)
	case *ast.GenerateStatement:
		e.Packages = append(e.Packages, e.evalGenerateStatement(node))
	}
	return err
}

func (e *Evaluator) evalStatements(stmts []ast.Statement) error {
	for _, statement := range stmts {
		err := e.eval(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Evaluator) evalSourceStatement(stmt *ast.SourceStatement) (source.Source, error) {
	return source.NewRemoteGitSource(e.logger, stmt.Source.String())
}

func (e *Evaluator) evalLanguageStatement(stmt *ast.LanguageStatement) config.Language {
	return config.Language(stmt.Name.String())
}

func (e *Evaluator) evalOutputStatement(stmt *ast.OutputStatement) string {
	return stmt.Path.String()
}

func (e *Evaluator) evalGenerateStatement(stmt *ast.GenerateStatement) *config.Package {
	var ref source.Ref
	switch t := stmt.Tag.(type) {
	case *ast.Version:
		strs := []string{stmt.Target.String(), t.String()}
		tagName := strings.Join(strs, "/")
		ref = source.Ref{
			Name: tagName,
			Type: source.Version,
		}
	default:
		ref = source.Ref{
			Name: t.String(),
			Type: source.Branch,
		}
	}

	return &config.Package{
		Source:   e.CurrentSource,
		Language: e.CurrentLanguage,
		Output:   e.CurrentOutput,
		Ref:      ref,
		Name:     stmt.Target.String(),
	}
}
