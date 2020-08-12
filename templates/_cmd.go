package cmd

import (
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/settings"
)

func Create{{.Command.ApiName}}Command(api api.Api) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.Command.Name}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			settings := settings.GetSettings() // Get the default settings
			return api.{{.Command.ApiName}}(settings) // Call api
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.Name}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
