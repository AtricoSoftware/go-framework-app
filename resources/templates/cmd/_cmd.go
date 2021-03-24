{"Type":"Framework", "Name":"%s"}
// {{.Comment}}
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

type {{.Command.ApiName}}Cmd *cobra.Command

func RegisterCmd{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func() {{.Command.ApiName}}Cmd { return create{{.Command.ApiName}}Command(c) })
}

func create{{.Command.ApiName}}Command(c container.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.Command.UseName}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			api.RegisterApi{{.Command.ApiName}}(c)
			var theApi api.{{.Command.ApiName}}Api
			c.Make(&theApi)
			return theApi.Run()
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
