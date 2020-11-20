package settings

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
{{- end}}