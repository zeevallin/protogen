package evaluator

import (
	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/source"
)

// New returns a new evaluation instance
func New() *Evaluator {
	return &Evaluator{
		Projects: make([]*config.Project, 0),
	}
}

// Evaluator evaluates a config node
type Evaluator struct {
	Projects        []*config.Project
	CurrentLanguage config.Language
	CurrentSource   source.Source
}

// Eval evaluates a configuration file AST
func (e *Evaluator) Eval(node *ast.ConfigurationFile) *config.Config {
	e.eval(node)
	return &config.Config{
		Projects: e.Projects,
	}
}

func (e *Evaluator) eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.ConfigurationFile:
		e.evalStatements(node.Statements)
	case *ast.SourceStatement:
		e.CurrentSource = e.evalSourceStatement(node)
	case *ast.LanguageStatement:
		e.CurrentLanguage = e.evalLanguageStatement(node)
	case *ast.GenerateStatement:
		e.Projects = append(e.Projects, e.evalGenerateStatement(node))
	}
	return
}

func (e *Evaluator) evalStatements(stmts []ast.Statement) {
	for _, statement := range stmts {
		e.eval(statement)
	}
	return
}

func (e *Evaluator) evalSourceStatement(stmt *ast.SourceStatement) source.Source {
	return &source.RemoteGitSource{
		URL: stmt.Source.String(),
	}
}

func (e *Evaluator) evalLanguageStatement(stmt *ast.LanguageStatement) config.Language {
	return config.Language(stmt.Name.String())
}

func (e *Evaluator) evalGenerateStatement(stmt *ast.GenerateStatement) *config.Project {
	return &config.Project{
		Source:   e.CurrentSource,
		Language: e.CurrentLanguage,
		Tag:      stmt.Tag.String(),
		Target:   stmt.Target.String(),
	}
}
