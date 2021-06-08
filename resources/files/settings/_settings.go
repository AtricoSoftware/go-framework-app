// {{.Comment}}
package settings

import "github.com/atrico-go/container"

type Settings interface {
	// Cmd line arguments
	GetArgument(name string) (value string, ok bool)
	MustGetArgument(name string) (value string)
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

func (s theSettings) GetArgument(name string) (value string, ok bool) {
	value, ok = s.args[name]
	return value, ok
}

func (s theSettings) MustGetArgument(name string) string {
	if value, ok := s.GetArgument("game"); ok {
		return value
	}
	panic(fmt.Sprintf("'%s' argument not found", name))
}
