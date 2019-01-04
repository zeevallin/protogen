package cli

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/zeeraw/protogen/source"

	"github.com/urfave/cli"
)

const (
	// Version is the package version
	Version = "0.0.1"

	// Name is the name of the command line tool
	Name = "protogen"

	// Usage is the description of the tool displayed in the cli help section
	Usage = "Command line tool for organising code generation of Google's protocol buffers across multiple projects"

	// UsageText is the general usage of the command line tool
	UsageText = "[OPTIONS] [COMMAND]"

	loggerTag = "protogen "
)

// Runner holds the configuration for the running context
type Runner struct {
	protogenFile string
	verbose      bool
	workDir      string
	lang         *string
	source       *string
}

// Run will
func Run(args []string) error {
	runner := &Runner{}
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Usage
	app.UsageText = UsageText
	app.Version = Version
	app.HideVersion = true
	app.Authors = authors
	app.Commands = []cli.Command{
		runner.cmdGenerate(),
		runner.cmdInit(),
		runner.cmdInfo(),
		runner.cmdVersion(),
	}
	app.Flags = []cli.Flag{
		runner.flagVerbose(),
		runner.flagWorkDir(),
	}
	app.Copyright = "Philip V. (Zee) – Apache 2.0"
	app.Before = func(ctx *cli.Context) error {
		// Setup the logger
		var logdest io.Writer
		if runner.verbose {
			logdest = os.Stdout
		} else {
			logdest = ioutil.Discard
		}
		log.SetOutput(logdest)
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.SetPrefix(loggerTag)

		// Set the workdir singleton of the source package
		source.WorkDir = runner.workDir
		return nil
	}
	return app.Run(args)
}
