{"Type":"Framework", "Name":"%s"}
// {{.Comment}}
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/api"
{{- $write := false}}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	{{- $write = true}}
	{{- end}}
{{- end}}
{{- if $write}}
	"{{.RepositoryPath}}/settings"
{{- end}}
)

type {{.Command.ApiName}}Cmd commandInfo

func RegisterCmd{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(apiFactory api.{{.Command.ApiName}}ApiFactory) {{.Command.ApiName}}Cmd { return {{.Command.ApiName}}Cmd(create{{.Command.ApiName}}Command(apiFactory)) })
}

func create{{.Command.ApiName}}Command(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "{{.Command.UseName}}",
		Short: "{{.Command.Description}}",
		RunE: func(*cobra.Command, []string) error {
			theApi := apiFactory.Create()
			return theApi.Run()
		},
	}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return commandInfo{cmd, "{{.Command.Name}}" }
}
