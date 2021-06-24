// {{.Comment}}
package unit_tests

import (
	"{{.RepositoryPath}}/settings"
)

type MockSettings struct {
{{- $table := createTable }}
{{- $table = $table.AppendRow "\tTheCommand" "[]string" }}
{{- $table = $table.AppendRow "\tTheArgs" "[]string" }}
{{- range .UserSettings}}
{{- $table = $table.AppendRow (print "\t" .NameCode "Var") .QualifiedType }}
{{- end}}
{{ printTable $table -}}
}
{{- range .UserSettings}}

func (s MockSettings) {{.NameCode}}() {{.QualifiedType}} {
	return s.{{.NameCode}}Var
}
{{- end}}

func NewMockSettings(cmd []string, args []string) MockSettings {
{{- $table := createTable }}
{{- $table = $table.AppendRow "\t\tTheCommand:" "cmd," }}
{{- $table = $table.AppendRow "\t\tTheArgs:" "args," }}
{{- range .UserSettings}}
{{- if (ne .DefaultVal "")}}
	{{- $quote := "" }}
	{{- if (eq .Type "string")}}{{ $quote = "\"" }}{{end}}
	{{- $table = $table.AppendRow (print "\t\t" .NameCode "Var:") (print $quote .DefaultVal $quote ",") }}
{{- else}}
{{- if (ne .InitCode "")}}
	{{- $table = $table.AppendRow (print "\t\t" .NameCode "Var:") (print .QualifiedInitCode ",") }}
{{- end}}
{{- end}}
{{- end}}
	return MockSettings{
{{ printTable $table -}}
	{{ print "\t}" }}
}