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
func ReadConfigFromFilePath(path string) (*config.Config, error) {
	log.Printf("reading file at %s\n", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("cannot read file: %s\n", err)
		return nil, err
	}
	l := lexer.New(f)
	p := parser.New(l)

	cfg, err := p.Parse()
	if err != nil {
		log.Printf("cannot parse file: %s\n", err)
		return nil, err
	}

	e := evaluator.New()
	return e.Eval(cfg)
}
