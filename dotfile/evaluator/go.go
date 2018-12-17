package evaluator

import (
	"fmt"

	"github.com/zeeraw/protogen/config/go"
	"github.com/zeeraw/protogen/dotfile/ast"
)

func (e *Evaluator) evalLanguageGoConfigBlock(blk *ast.Block) (golang.Config, error) {
	var (
		paths   = golang.SourceRelative
		plugins = []golang.Plugin{}
		cfg     = golang.Config{}
	)
	for _, stmt := range blk.Statements {
		var err error
		switch node := stmt.(type) {
		case *ast.GoPathStatement:
			paths, err = e.evalGoPath(node)
			if err != nil {
				return cfg, err
			}
		case *ast.GoPluginStatement:
			plugin, err := e.evalGoPlugin(node)
			if err != nil {
				return cfg, err
			}
			plugins = append(plugins, plugin)
		default:
			return cfg, fmt.Errorf("format string")
		}
	}
	cfg.Paths = paths
	cfg.Plugins = plugins
	return cfg, nil
}

func (e *Evaluator) evalGoPlugin(stmt *ast.GoPluginStatement) (golang.Plugin, error) {
	plugin := golang.Plugin(stmt.Name.String())
	return plugin, golang.IsAllowedPlugin(plugin)
}
func (e *Evaluator) evalGoPath(stmt *ast.GoPathStatement) (golang.Path, error) {
	path := golang.Path(stmt.Type.String())
	return path, golang.IsAllowedPath(path)
}
