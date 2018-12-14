package cli

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/zeeraw/protogen/generator/protoc"
)

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
	version, err := protoc.NewProtoc(r.logger()).Check()
	if err != nil {
		return err
	}
	fmt.Println(version)
	return nil
}
