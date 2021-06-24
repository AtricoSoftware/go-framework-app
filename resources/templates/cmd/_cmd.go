{"Type":"Framework", "Name":"%s"}
// {{.Comment}}
package cmd

import (
{{- if not .Command.NoImplementation}}
	"fmt"
	"os"
{{- end}}
	"github.com/atrico-go/container"
{{- if .Command.HasArgs}}
	"github.com/atrico-go/cobraEx"
{{- end}}
	"github.com/spf13/cobra"

{{- if not .Command.NoImplementation}}
	"{{.RepositoryPath}}/api"
{{- end}}
{{- $write := false}}
{{- if not .Command.NoImplementation}}
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	{{- $write = true}}
	{{- end}}
{{- end}}
{{- end}}
{{- if $write}}
	"{{.RepositoryPath}}/settings"
{{- end}}
)

type {{.Command.ApiName}}Cmd commandInfo

func RegisterCmd{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func({{ if not .Command.NoImplementation}}apiFactory api.{{.Command.ApiName}}ApiFactory{{end}}) {{.Command.ApiName}}Cmd { return {{.Command.ApiName}}Cmd(create{{.Command.ApiName}}Command({{ if not .Command.NoImplementation}}apiFactory{{end}})) })
}

func create{{.Command.ApiName}}Command({{ if not .Command.NoImplementation}}apiFactory api.Factory{{end}}) commandInfo {
	cmd := &cobra.Command{
		Use:   "{{.Command.UseName}}",
		Short: "{{.Command.Description}}",
{{- if not .Command.NoImplementation}}
		Args: cobra.{{.Command.ArgsValidator}},
		Run: func(cmd *cobra.Command, args []string) {
			theApi := apiFactory.Create(args)
			if err := theApi.Run(); err != nil {
				fmt.Fprint(os.Stderr, err)
			}
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
{{- else}}
	}
{{- end}}
	return commandInfo{cmd, "{{.Command.Name}}" }
}
