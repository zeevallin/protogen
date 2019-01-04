package evaluator

import (
	"log"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/swift"
	"github.com/zeeraw/protogen/dotfile/ast"
)

func (e *Evaluator) evalLanguageSwiftConfigBlock(blk *ast.Block) (swift.Config, error) {
	log.Printf("evaluating swift config block with %d statements\n", len(blk.Statements))
	var (
		cfg     = swift.Config{}
		options []swift.Option
	)
	for _, stmt := range blk.Statements {
		switch node := stmt.(type) {
		case *ast.OptionStatement:
			option, err := e.evalSwiftOption(node)
			if err != nil {
				return cfg, err
			}
			options = append(options, option)
		default:
			return cfg, ErrStatementNotSupported{config.Swift, stmt}
		}
	}
	cfg.Options = options
	return cfg, nil
}

func (e *Evaluator) evalSwiftOption(stmt *ast.OptionStatement) (swift.Option, error) {
	return swift.Option{
		Name:  stmt.Name.String(),
		Value: stmt.Value.String(),
	}, nil
}
