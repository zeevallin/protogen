package main

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/zeeraw/protogen/dotfile/evaluator"
	"github.com/zeeraw/protogen/dotfile/lexer"
	"github.com/zeeraw/protogen/dotfile/parser"
)

const (
	// VERSION is the package version
	VERSION = "1.0.0"

	// Name is the name of the command line tool
	Name = "protogen"

	// Usage is the description of the tool displayed in the cli help section
	Usage = "Helps generate and organise your code from .proto files"
)

var (
	authorPhilipV = cli.Author{
		Name:  "Philip V. (Zee)",
		Email: "zee@vall.in",
	}
)

// Runner holds the configuration for the running context
type Runner struct {
	configFilePath *string
}

func main() {
	runner := &Runner{}
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Usage
	app.Version = VERSION
	app.Authors = []cli.Author{authorPhilipV}
	app.Action = action
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "protogen_file",
			Usage:       "path to the protogen configuration file for the current project",
			Destination: runner.configFilePath,
			Value:       "./.protogen",
		},
	}
	app.Run(os.Args)
}

func action(cc *cli.Context) error {
	_, cancel := makeCtx()
	defer cancel()

	f, err := ioutil.ReadFile(".protogen")
	if err != nil {
		panic(err)
	}

	l := lexer.New(f)
	p := parser.New(l)

	cf, err := p.ParseConfigurationFile()
	if err != nil {
		panic(err)
	}

	e := evaluator.New()
	e.Eval(cf)

	return nil
}

func makeCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*30)
}
