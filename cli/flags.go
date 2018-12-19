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

func (r *Runner) flagWorkDir() cli.Flag {
	return &cli.StringFlag{
		Name:        "work_dir",
		Usage:       "the path on disk where protogen operates and saves its cache",
		Destination: &r.workDir,
		Value:       DefaultWorkDir,
		EnvVar:      "PROTOGEN_WORKDIR",
	}
}

func (r *Runner) flagProtogenFile() cli.Flag {
	return &cli.StringFlag{
		Name:        "protogen_file",
		Usage:       "path to the protogen configuration file for the current project",
		Destination: &r.protogenFile,
		Value:       ".protogen",
		EnvVar:      "PROTOGEN_FILE",
	}
}

func (r *Runner) flagSource() cli.Flag {
	return &cli.StringFlag{
		Name:        "source",
		Usage:       "the remote repository for your proto files",
		Destination: r.source,
	}
}

func (r *Runner) flagLang() cli.Flag {
	return &cli.StringFlag{
		Name:        "lang",
		Usage:       "the language you want to generate code for in this project",
		Destination: r.lang,
	}
}
