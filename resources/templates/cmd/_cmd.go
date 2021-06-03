{"Type":"Framework", "Name":"%s"}
// {{.Comment}}
package cmd

import (
	"github.com/atrico-go/container"
{{- if .Command.HasArgs}}
	"github.com/atrico-go/cobraEx"
{{- end}}
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
		Args: cobra.{{.Command.ArgsValidator}},
		RunE: func(cmd *cobra.Command, args []string) error {
			theApi := apiFactory.Create()
			return theApi.Run(args)
		},
	}
{{- if .Command.HasArgs}}
	cobraEx.AddUsageParameters(cmd, []string{ {{- commaList (quoted .Command.Args) -}} }, []string{ {{- commaList (quoted .Command.OptionalArgs) -}} })
{{- end}}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return commandInfo{cmd, "{{.Command.Name}}" }
}
