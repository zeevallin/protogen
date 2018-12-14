package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	authorPhilipV = cli.Author{
		Name:  "Philip V. (Zee)",
		Email: "zee@vall.in",
	}
	authors = []cli.Author{
		authorPhilipV,
	}
)

func (r *Runner) cmdAuthors() cli.Command {
	const usage = "shows a list of the project authors"
	const description = "Shows a list of the project authors."
	return cli.Command{
		Name:        "authors",
		Usage:       usage,
		Description: description,
		Action:      r.authors,
	}
}

func (r *Runner) authors(cc *cli.Context) error {
	for _, author := range authors {
		fmt.Println(author.String())
	}
	return nil
}
