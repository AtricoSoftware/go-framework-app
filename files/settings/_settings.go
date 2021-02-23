package settings

import "github.com/atrico-go/container"

type Settings interface {
{{- range .UserSettings}}
	// {{.Description}}
	{{.NameCode}}() {{.Type}}
{{- end}}
}

// Register the settings
func RegisterSettings(c container.Container) {
	c.Singleton(func() Settings {return theSettings{}})
}

// Stub object for settings interface
type theSettings struct{}
