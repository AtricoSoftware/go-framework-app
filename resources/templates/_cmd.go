package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/api"
{{- $write := false}}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.UseName}}
	{{- $write = true}}
	{{- end}}
{{- end}}
{{- if $write}}
	"{{.RepositoryPath}}/settings"
{{- end}}
)

func Create{{.Command.ApiName}}Command(c container.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.Command.UseName}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			var {{.Command.Name}}Api api.{{.Command.ApiName}}Api
			c.Make(&{{.Command.Name}}Api)
			return {{.Command.Name}}Api.Run()
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.UseName}}
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}