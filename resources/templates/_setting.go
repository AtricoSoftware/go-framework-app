package settings

{{- $settingVarName := print .Setting.LowerName "SettingName"}}
{{- $cmdlineVarName := print .Setting.LowerName "SettingCmdline"}}
{{- $shortcutVarName := print .Setting.LowerName "SettingShortcut"}}
{{- $defaultVarName := print .Setting.LowerName "SettingDefaultVal"}}
{{- $lazyVarName := print .Setting.LowerName "SettingLazy"}}

{{- $getTheValue := ""}}
{{- if (ne .Setting.TypeGetter "")}}
{{- $getTheValue = print .Setting.TypeGetter "(" $settingVarName ")"}}
{{- else}}
{{- $getTheValue = print "Parse" .Setting.NameCode "Setting(viper.Get(" $settingVarName "))"}}
{{- end}}

import (
{{- if or (and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")) (.Setting.HasPrefix .Setting.TypeGetter "viperEx.")}}
	"github.com/atrico-go/viperEx"
{{- end}}
{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}
	"github.com/spf13/pflag"
{{- end}}
{{- if or (.Setting.HasPrefix .Setting.TypeGetter "viper.") (eq .Setting.TypeGetter "")}}
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

{{- if .SingleReadConfiguration}}

// Lazy value
var {{$lazyVarName}} = NewLazy{{.Setting.TypeNameAsCode}}Value(func() {{.Setting.Type}} { return {{$getTheValue}} })
{{- end}}

// Fetch the setting
func (theSettings) {{.Setting.NameCode}}() {{.Setting.Type}} {
{{- if .SingleReadConfiguration}}
	return {{$lazyVarName}}.GetValue()
{{- else}}
	return {{$getTheValue}})
{{- end}}
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
