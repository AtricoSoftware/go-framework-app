
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
	settings.Add{{.NameCode}}Flag(cmd.PersistentFlags())
	{{- end}}
{{- end}}
	return cmd
}
`))
Templates[`setting`] = template.Must(template.New(`setting`).Parse(`package settings

{{- $settingVarName := print .Setting.LowerName "SettingName"}}
{{- $cmdlineVarName := print .Setting.LowerName "SettingCmdline"}}
{{- $shortcutVarName := print .Setting.LowerName "SettingShortcut"}}
{{- $defaultVarName := print .Setting.LowerName "SettingDefaultVal"}}

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

const {{$settingVarName}} = "{{.Setting.Id}}"
{{- if (ne .Setting.Cmdline "")}}
const {{$cmdlineVarName}} = "{{.Setting.Cmdline}}"
{{- end}}
{{- if (ne .Setting.CmdlineShortcut "")}}
const {{$shortcutVarName}} = "{{.Setting.CmdlineShortcut}}"
{{- end}}
{{- if (ne .Setting.DefaultVal "")}}
const {{$defaultVarName}} = {{if (eq .Setting.Type "string")}}"{{end}}{{.Setting.DefaultVal}}{{if (eq .Setting.Type "string")}}"{{end}}
{{- end}}

// Fetch the setting
func (theSettings) {{.Setting.NameCode}}() {{.Setting.Type}} {
	return {{.Setting.TypeGetter}}({{$settingVarName}})
}

{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}

func Add{{.Setting.NameCode}}Flag(flagSet *pflag.FlagSet) {
	{{.Setting.TypeFlagAdder}}{{if (ne .Setting.CmdlineShortcut "")}}P{{end}}(flagSet, {{$settingVarName}}, {{$cmdlineVarName}}, {{if (ne .Setting.CmdlineShortcut "")}}{{$shortcutVarName}}, {{end}}"{{.Setting.Description}}")
}
{{- end}}
{{- if (ne .Setting.DefaultVal "")}}

func init() {
	viper.SetDefault({{$settingVarName}}, {{$defaultVarName}})
}
{{- end}}
`))

}