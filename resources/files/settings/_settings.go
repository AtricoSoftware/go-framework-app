// {{.Comment}}
package settings

import "github.com/atrico-go/container"

type Settings interface {
	// Cmd line arguments
	Argument(name string) (value string, ok bool)
{{- range .UserSettings}}
	// {{.Description}}
	{{.NameCode}}() {{.Type}}
{{- end}}
}

type SetArgs interface {
	SetArgs(map[string]string)
}

// Register the settings
func RegisterSettings(c container.Container) {
	settings := theSettings{make(map[string]string)}
	c.Singleton(func() Settings { return &settings })
	c.Singleton(func() SetArgs { return &settings })
}
{{- if .SingleReadConfiguration}}

// Force all settings to be recalculated on next request
func ResetCaches() {
{{- range .UserSettings}}
	{{.LowerName}}SettingCache.Reset()
{{- end}}
}
{{- end}}

// Stub object for settings interface
type theSettings struct{
	args map[string]string
}

func (s *theSettings) SetArgs(args map[string]string) {
	s.args = args
}

func (s theSettings) Argument(name string) (value string, ok bool) {
	value, ok = s.args[name]
	return value, ok
}