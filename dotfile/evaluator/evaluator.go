package evaluator

import (
	"log"
	"strings"

	"github.com/zeeraw/protogen/config/go"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/dotfile/ast"
	"github.com/zeeraw/protogen/source"
)

// New returns a new evaluation instance
func New() *Evaluator {
	return &Evaluator{
		Packages: make([]*config.Package, 0),
	}
}

// Evaluator evaluates a config node
type Evaluator struct {
	Packages              []*config.Package
	CurrentLanguage       config.Language
	CurrentLanguageConfig interface{}
	CurrentSource         source.Source
	CurrentOutput         string
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
		e.CurrentLanguageConfig, err = e.evalLanguageConfigStatement(node)
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
	src := stmt.Source.String()
	log.Printf("evaluating source statement: %s\n", src)
	if isFilePath(src) {
		return source.NewLocalGitSource(src)
	}
	return source.NewRemoteGitSource(src)
}

func isFilePath(s string) bool {
	return strings.HasPrefix(s, "./") || strings.HasPrefix(s, "~/") || strings.HasPrefix(s, "/")
}

func (e *Evaluator) evalLanguageStatement(stmt *ast.LanguageStatement) config.Language {
	log.Printf("evaluating language statement: %s\n", stmt.Name.String())
	return config.Language(stmt.Name.String())
}

func (e *Evaluator) evalLanguageConfigStatement(stmt *ast.LanguageStatement) (interface{}, error) {
	log.Printf("evaluating language config: %s\n", stmt.Name.String())
	lang := config.Language(stmt.Name.String())
	switch lang {
	case config.Go:
		if stmt.Block != nil {
			return e.evalLanguageGoConfigBlock(stmt.Block)
		}
		return golang.Config{}, nil
	default:
		if stmt.Block != nil {
			err := ErrLanguageNotSupported{lang}
			log.Printf("cannot evaluate unsupported: %v\n", err)
			return nil, err
		}
		return nil, nil
	}
}

func (e *Evaluator) evalOutputStatement(stmt *ast.OutputStatement) string {
	log.Printf("evaluating output statement: %s\n", stmt.Path.String())
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

	log.Printf("evaluating generate statement: %q %T(%s)\n", stmt.Target.String(), ref.Type, ref.Name)

	return &config.Package{
		Source:         e.CurrentSource,
		Language:       e.CurrentLanguage,
		LanguageConfig: e.CurrentLanguageConfig,
		Output:         e.CurrentOutput,
		Ref:            ref,
		Name:           stmt.Target.String(),
	}
}
