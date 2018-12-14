package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

func (r *Runner) cmdVersion() cli.Command {
	const usage = "prints the version of the command line tool"
	const description = `Prints the tagged version of the command line tool`
	return cli.Command{
		Name:        "version",
		Usage:       usage,
		Description: description,
		Action:      r.version,
	}
}

func (r *Runner) version(cc *cli.Context) error {
	fmt.Println(Version)
	return nil
}
