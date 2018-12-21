package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/generator/protoc"
)

const (
	checkMark          = "✓"
	crossMark          = "✘"
	fmtExtBin          = "protoc-gen-%s"
	fmtCheckBin        = "%s\t %s\n"
	fmtCheckBinVersion = "%s\t %s\t (%s)\n"
	fmtAuthor          = "%s\t %s\n"
)

var languages = []config.Language{
	config.Go,
	config.Swift,
}

var (
	authorPhilipV = cli.Author{
		Name:  "Philip V. (Zee)",
		Email: "zee@vall.in",
	}
	authors = []cli.Author{
		authorPhilipV,
	}
)

func (r *Runner) cmdInfo() cli.Command {
	const usage = "shows information about protogen and its dependencies"
	const description = `Gathers and displays information about the protogen environment and its dependencies.`
	return cli.Command{
		Name:        "info",
		Usage:       usage,
		Description: description,
		Action:      r.info,
	}
}

func (r *Runner) info(cc *cli.Context) error {
	fmt.Printf("Dependencies\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	p := protoc.NewProtoc(r.logger())
	version, err := p.Check()
	if err != nil {
		fmt.Fprintf(w, fmtCheckBin, p.Binary, crossMark)
	} else {
		fmt.Fprintf(w, fmtCheckBinVersion, p.Binary, checkMark, version)
	}

	for _, lang := range languages {
		extBin := fmt.Sprintf(fmtExtBin, lang)
		version, err := p.CheckExtension(lang)
		if err != nil {
			fmt.Fprintf(w, fmtCheckBin, extBin, crossMark)
			continue
		}
		if version == "" {
			fmt.Fprintf(w, fmtCheckBin, extBin, checkMark)
			continue
		}
		fmt.Fprintf(w, fmtCheckBinVersion, extBin, checkMark, version)
	}
	w.Flush()

	fmt.Printf("\nAuthors\n")
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	for _, author := range authors {
		fmt.Fprintf(w, fmtAuthor, author.Name, author.Email)
	}
	w.Flush()

	return nil
}
