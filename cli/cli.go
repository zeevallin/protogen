package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/generator"

	"github.com/urfave/cli"
)

const (
	// VERSION is the package version
	VERSION = "0.0.1"

	// Name is the name of the command line tool
	Name = "protogen"

	// Usage is the description of the tool displayed in the cli help section
	Usage = "Helps generate and organise your code from .proto files"

	loggerTag = "protogen "
)

var runner = &Runner{}

var (
	authorPhilipV = cli.Author{
		Name:  "Philip V. (Zee)",
		Email: "zee@vall.in",
	}
	authors = []cli.Author{
		authorPhilipV,
	}

	flagVerbose = cli.BoolFlag{
		Name:        "verbose",
		Usage:       "use this to show more information",
		Destination: &runner.verbose,
	}
	flagProtogenFile = cli.StringFlag{
		Name:        "protogen_file",
		Usage:       "path to the protogen configuration file for the current project",
		Destination: &runner.configFilePath,
		Value:       "./.protogen",
		EnvVar:      "PROTOGEN_FILE",
	}
)

// Runner holds the configuration for the running context
type Runner struct {
	configFilePath string
	verbose        bool
}

// Run will
func Run(args []string) error {
	runner = &Runner{}
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Usage
	app.Version = VERSION
	app.Authors = authors
	app.Action = action
	app.Flags = []cli.Flag{
		flagProtogenFile,
		flagVerbose,
	}
	return app.Run(args)
}

func action(cc *cli.Context) error {
	l := logger()
	cfg, err := ReadConfigFromFilePath(l, runner.configFilePath)
	if err != nil {
		return err
	}
	cfg.General = &config.General{
		Verbose: runner.verbose,
	}
	if err := generator.Generate(l, cfg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func logger() *log.Logger {
	var logdest io.Writer
	if runner.verbose {
		logdest = os.Stdout
	} else {
		logdest = ioutil.Discard
	}
	return log.New(logdest, loggerTag, log.Lshortfile|log.LstdFlags)
}
