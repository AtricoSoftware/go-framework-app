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
	c.Singleton(func(api api.{{.Command.ApiName}}Api) {{.Command.ApiName}}Cmd { return create{{.Command.ApiName}}Command(api) })
}

func create{{.Command.ApiName}}Command(api api.Runnable) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{.Command.UseName}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			return api.Run()
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
