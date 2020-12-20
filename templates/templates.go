
package templates

import "text/template"

// Specific file templates
var Templates = make(map[string]*template.Template)

func init() {
Templates[`api`] = template.Must(template.New(`api`).Parse(`package api

import (
	"github.com/atrico-go/container"

  	"{{.RepositoryPath}}/settings"
)

func Register{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings) {{.Command.ApiName}}Api {return {{.Command.Name}}Api{config: config}})
}

type {{.Command.Name}}Api struct {
config settings.Settings
}

// {{.Command.Description}}
func (svc {{.Command.Name}}Api) Run() error {
	// Implementation here!
	return nil
}
`))
Templates[`cmd`] = template.Must(template.New(`cmd`).Parse(`package cmd

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
	settings.Add{{.Name}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
`))
Templates[`setting`] = template.Must(template.New(`setting`).Parse(`package settings

import (
{{- if or (and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")) (.Setting.HasPrefix .Setting.TypeGetter "viperEx.")}}
"github.com/atrico-go/viperEx"
{{- end}}
{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}
	"github.com/spf13/pflag"
{{- end}}
{{- if .Setting.HasPrefix .Setting.TypeGetter "viper."}}
	"github.com/spf13/viper"
{{- end}}
)

const {{.Setting.Id}}SettingName = "{{.Setting.Cmdline}}"
{{- if (ne .Setting.CmdlineShortcut "")}}
const {{.Setting.Id}}SettingShortcut = "{{.Setting.CmdlineShortcut}}"
{{- end}}
{{- if (ne .Setting.DefaultVal "")}}
const {{.Setting.Id}}SettingDefaultVal = {{if (eq .Setting.Type "string")}}"{{end}}{{.Setting.DefaultVal}}{{if (eq .Setting.Type "string")}}"{{end}}
{{- end}}

// Fetch the setting
func (theSettings) {{.Setting.Name}}() {{.Setting.Type}} {
	return {{.Setting.TypeGetter}}({{.Setting.Id}}SettingName)
}

{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}

func Add{{.Setting.Name}}Flag(flagSet *pflag.FlagSet) {
	{{.Setting.TypeFlagAdder}}{{if (ne .Setting.CmdlineShortcut "")}}P{{end}}{{if (ne .Setting.DefaultVal "")}}D{{end}}(flagSet, {{.Setting.Id}}SettingName, {{if (ne .Setting.CmdlineShortcut "")}}{{.Setting.Id}}SettingShortcut, {{end}}{{if (ne .Setting.DefaultVal "")}}{{.Setting.Id}}SettingDefaultVal, {{end}}"{{.Setting.Description}}")
}
{{- end}}`))

}