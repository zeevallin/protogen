package cli

import (
	"io"
	"io/ioutil"
	"log"
	"os"

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

var (
	// GitCommit is the git commit hash of the source-tree during build
	GitCommit string
)

// Runner holds the configuration for the running context
type Runner struct {
	protogenFile string
	verbose      bool
	lang         *string
	source       *string
}

func (r *Runner) logger() *log.Logger {
	var logdest io.Writer
	if r.verbose {
		logdest = os.Stdout
	} else {
		logdest = ioutil.Discard
	}
	return log.New(logdest, loggerTag, log.Lshortfile|log.LstdFlags)
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
		runner.cmdCheck(),
		runner.cmdAuthors(),
		runner.cmdVersion(),
	}
	app.Flags = []cli.Flag{
		runner.flagVerbose(),
	}
	app.Copyright = "Philip V. (Zee) – Apache 2.0"

	return app.Run(args)
}
