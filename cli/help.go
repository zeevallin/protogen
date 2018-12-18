package cli

import (
	"github.com/urfave/cli"
)

func init() {
	cli.AppHelpTemplate = `Usage: {{.Name}} {{if .UsageText}}{{.UsageText}}{{end}}
{{if .VisibleCommands}}
Commands:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
	 {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
{{if .VisibleFlags}}
Options:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}

Run '{{.Name}} help COMMAND' for more information on a command.
`

	cli.CommandHelpTemplate = `Usage: {{.HelpName}}{{if .VisibleFlags}} [OPTIONS]{{end}}{{if .ArgsUsage}} {{.ArgsUsage}}{{else}} [ARGUMENTS...]{{end}}

{{.Description}}
{{if .VisibleFlags}}
Options:
	{{range .VisibleFlags}}{{.}}
	{{end -}}
{{end -}}
`
}
