package cli

import (
	"io/ioutil"
	"log"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/dotfile/evaluator"
	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/parser"
)

// ReadConfigFromFilePath will attempt to read a .protogen configuration file at the given path
func ReadConfigFromFilePath(logger *log.Logger, path string) (*config.Config, error) {
	logger.Printf("reading file at %s\n", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	l := lexer.New(logger, f)
	p := parser.New(logger, l)

	cfg, err := p.Parse()
	if err != nil {
		return nil, err
	}

	e := evaluator.New(logger)
	return e.Eval(cfg)
}
