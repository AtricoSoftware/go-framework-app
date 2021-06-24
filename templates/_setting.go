package settings

import (
{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}
	"github.com/spf13/pflag"
{{- end}}
{{- if .Setting.HasPrefix .Setting.TypeGetter "viper."}}
	"github.com/spf13/viper"
{{- end}}

{{- if or (and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")) (.Setting.HasPrefix .Setting.TypeGetter "viperEx.")}}
	"{{.RepositoryPath}}/viperEx"
{{- end}}
)

const {{.Setting.Id}}SettingName = "{{.Setting.Id}}"

// Fetch the setting
func (theSettings) {{.Setting.Name}}() {{.Setting.Type}} {
	return {{.Setting.TypeGetter}}("{{.Setting.Id}}SettingName")
}

{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}

func Add{{.Setting.Name}}Flag(flagSet *pflag.FlagSet) {
	{{.Setting.TypeFlagAdder}}{{if (ne .Setting.CmdlineShortcut "")}}P{{end}}(flagSet, "{{.Setting.Id}}SettingName", {{if (ne .Setting.CmdlineShortcut "")}}"{{.Setting.CmdlineShortcut}}", {{end}}"{{.Setting.Description}}")
}
{{- end}}