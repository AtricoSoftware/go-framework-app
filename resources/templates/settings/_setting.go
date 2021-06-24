{"Type":"Mixed", "Name":"%s"}
// {{.Comment}}
// SECTION-START: Framework
package settings

{{- $settingVarName := print .Setting.LowerName "SettingName"}}
{{- $cmdlineVarName := print .Setting.LowerName "SettingCmdline"}}
{{- $shortcutVarName := print .Setting.LowerName "SettingShortcut"}}
{{- $envVarName := print .Setting.LowerName "SettingEnvVar"}}
{{- $defaultVarName := print .Setting.LowerName "SettingDefaultVal"}}
{{- $cachedVarName := print .Setting.LowerName "SettingCache"}}

{{- $getTheValue := ""}}
{{- if (ne .Setting.TypeGetter "")}}
{{- $getTheValue = print .Setting.TypeGetter "(" $settingVarName ")"}}
{{- else}}
{{- $getTheValue = print "Parse" .Setting.NameCode "Setting(viper.Get(" $settingVarName "))"}}
{{- end}}

import (
{{- if or (and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")) (.Setting.HasPrefix .Setting.TypeGetter "viperEx.")}}
	"github.com/atrico-go/viperEx/v2"
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
const {{$shortcutVarName}} = '{{.Setting.CmdlineShortcut}}'
{{- end}}
{{- if (ne .Setting.EnvVar "")}}
const {{$envVarName}} = "{{.Setting.EnvVar}}"
{{- end}}
{{- if (ne .Setting.DefaultVal "")}}
const {{$defaultVarName}} = {{if (eq .Setting.Type "string")}}"{{end}}{{.Setting.DefaultVal}}{{if (eq .Setting.Type "string")}}"{{end}}
{{- end}}

{{- if .SingleReadConfiguration}}

// Cached value
var {{$cachedVarName}} = NewCached{{.Setting.TypeNameAsCode}}Value(func() {{.Setting.Type}} { return {{$getTheValue}} })
{{- end}}

// Fetch the setting
func (theSettings) {{.Setting.NameCode}}() {{.Setting.Type}} {
{{- if .SingleReadConfiguration}}
	return {{$cachedVarName}}.GetValue()
{{- else}}
	return {{$getTheValue}}
{{- end}}
}

{{- if and (gt (len .Setting.AppliesTo) 0) (ne .Setting.Cmdline "")}}

func Add{{.Setting.NameCode}}Flag(flagSet *pflag.FlagSet) {
	{{.Setting.TypeFlagAdder}}({{$settingVarName}}, "{{.Setting.Description}}"){{if (ne .Setting.Cmdline "")}}.Cmdline({{$cmdlineVarName}}){{end}}{{if (ne .Setting.CmdlineShortcut "")}}.CmdlineShortcut({{$shortcutVarName}}){{end}}{{if (ne .Setting.EnvVar "")}}.EnvVar({{$envVarName}}){{end}}{{if (ne .Setting.DefaultVal "")}}.DefaultVal({{$defaultVarName}}){{end}}.AddTo(flagSet)
}
{{- else}}
{{- if (ne .Setting.DefaultVal "")}}

func init() {
	viper.SetDefault({{$settingVarName}}, {{$defaultVarName}})
}
{{- end}}
{{- end}}
// SECTION-END
{{- if (eq .Setting.TypeGetter "")}}

func Parse{{.Setting.NameCode}}Setting(setting interface{}) {{.Setting.Type}} {
	// TODO - Implementation here
	return nil
}

type {{.Setting.RawType}} struct{}
{{- end}}
