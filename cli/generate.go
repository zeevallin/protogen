package cli

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/generator"
)

func (r *Runner) cmdGenerate() cli.Command {
	const usage = "performs code generation based on the protogen configuration"
	const description = `Your source will be retrieved, checked out and have protoc run for
the relevant files on every release tag you've specified in your
.protogen file`

	return cli.Command{
		Name:        "generate",
		Usage:       usage,
		Description: description,
		Action:      r.generate,
	}
}

func (r *Runner) generate(cc *cli.Context) error {
	cfg, err := ReadConfigFromFilePath(r.logger(), r.protogenFile)
	if err != nil {
		return err
	}
	cfg.General = &config.General{
		Verbose: r.verbose,
	}
	if err := generator.Generate(r.logger(), cfg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}