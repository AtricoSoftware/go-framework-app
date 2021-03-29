// {{.Comment}}
package unit_tests

type MockSettings struct {
	TheCommand	string
{{- range .UserSettings}}
	{{.NameCode}}Var {{.QualifiedType}}
{{- end}}
}
{{- range .UserSettings}}

	func (s MockSettings) {{.NameCode}}() {{.QualifiedType}} {
		return s.{{.NameCode}}Var
	}
{{- end}}

func NewMockSettings(cmd string) MockSettings {
	return MockSettings{
		TheCommand: cmd,
{{- range .UserSettings}}
{{- if (ne .InitCode "")}}
		{{.NameCode}}Var: {{.QualifiedInitCode}},
{{- end}}
{{- end}}
	}
}