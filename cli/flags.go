package cli

import (
	"github.com/urfave/cli"
)

func (r *Runner) flagVerbose() cli.Flag {
	return &cli.BoolFlag{
		Name:        "verbose",
		Usage:       "use this to show more information",
		Destination: &r.verbose,
	}
}

func (r *Runner) flagProtogenFile() cli.Flag {
	return &cli.StringFlag{
		Name:        "protogen_file",
		Usage:       "path to the protogen configuration file for the current project",
		Destination: &r.protogenFile,
		Value:       "./.protogen",
		EnvVar:      "PROTOGEN_FILE",
	}
}