package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/settings"
)

func Create{{.Command.ApiName}}Command(c container.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.Command.Name}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			var {{.Command.Name}}Api api.{{.Command.ApiName}}Api
			c.Make(&{{.Command.Name}}Api)
			return {{.Command.Name}}Api.Run()
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.Name}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
