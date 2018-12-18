package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zeeraw/protogen/dotfile/ast"

	"github.com/urfave/cli"
)

func (r *Runner) cmdInit() cli.Command {
	const name = "init"
	const usage = "setup protogen for your current repository"
	const description = `Takes you through the process of setting up protogen for your current repository.`
	return cli.Command{
		Name:        name,
		Usage:       usage,
		Description: description,
		Action:      r.init,
		Flags: []cli.Flag{
			r.flagSource(),
			r.flagLang(),
		},
	}
}

func (r *Runner) init(cc *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)
	if r.source == nil {
		fmt.Print("Source repository: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		r.source = &text
	}
	if r.lang == nil {
		fmt.Print("Language to generate for in this project: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		r.lang = &text
	}

	f := ast.NewConfigurationFile()
	f.Statements = append(f.Statements, ast.NewSourceStatement(*r.source))
	f.Statements = append(f.Statements, ast.NewLanguageStatement(*r.lang))

	fo, err := os.Create(".protogen")
	if err != nil {
		return err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fo.WriteString(f.String())
	fo.WriteString("\n")

	return nil
}
