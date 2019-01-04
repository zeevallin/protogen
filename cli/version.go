package cli

import (
	"fmt"
	"strings"

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
	b := strings.Builder{}
	b.WriteString(Version)
	if GitCommit != "" {
		commit := string([]rune(GitCommit)[0:12])
		b.WriteString(fmt.Sprintf(" (%s)", commit))
	}
	fmt.Println(b.String())
	return nil
}
