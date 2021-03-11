// {{.Comment}}
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
	c.Singleton(func() Settings { return theSettings{} })
}

// Force all settings to be recalculated on next request
func ResetCaches() {
{{- range .UserSettings}}
	{{.LowerName}}SettingCache.Reset()
{{- end}}
}

// Stub object for settings interface
type theSettings struct{}
