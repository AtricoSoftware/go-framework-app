// {{.Comment}}
package unit_tests

import (
	"{{.RepositoryPath}}/settings"
)

type MockSettings struct {
	TheCommand	[]string
	Args        []string
{{- range .UserSettings}}
	{{.NameCode}}Var {{.QualifiedType}}
{{- end}}
}
{{- range .UserSettings}}
func (s MockSettings) {{.NameCode}}() {{.QualifiedType}} {
	return s.{{.NameCode}}Var
}
{{- end}}

func NewMockSettings(cmd []string, args []string) MockSettings {
	return MockSettings{
		TheCommand: cmd,
		Args: args,
{{- range .UserSettings}}
{{- if (ne .DefaultVal "")}}
		{{.NameCode}}Var: {{if (eq .Type "string")}}"{{end}}{{.DefaultVal}}{{if (eq .Type "string")}}"{{end}},
{{- else}}
{{- if (ne .InitCode "")}}
		{{.NameCode}}Var: {{.QualifiedInitCode}},
{{- end}}
{{- end}}
{{- end}}
	}
}