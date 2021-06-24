package settings

type Settings interface {
{{- range .UserSettings}}
	// {{.Description}}
	{{.Name}}() {{.Type}}
{{- end}}
}

// Get the settings for this run
func GetSettings() Settings {
	return theSettings{}
}

// Stub object for settings interface
type theSettings struct{}
