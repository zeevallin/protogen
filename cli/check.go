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
)

var languages = []config.Language{
	config.Go,
}

func (r *Runner) cmdCheck() cli.Command {
	const usage = "checks verions of protogen dependencies"
	const description = `The filesystem will be scanned for all relevant dependencies and return a list
containing name and version of the dependency.`
	return cli.Command{
		Name:        "check",
		Usage:       usage,
		Description: description,
		Action:      r.check,
	}
}

func (r *Runner) check(cc *cli.Context) error {
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
		err := p.CheckExtension(lang)
		if err != nil {
			fmt.Fprintf(w, fmtCheckBin, extBin, crossMark)
		} else {
			fmt.Fprintf(w, fmtCheckBin, extBin, checkMark)
		}

	}
	w.Flush()

	return nil
}
