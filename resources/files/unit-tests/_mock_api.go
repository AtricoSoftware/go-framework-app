// {{.Comment}}
package unit_tests

import (
	"github.com/atrico-go/container"

	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/settings"
)

type mockApi struct {
	cmd    string
	config settings.Settings
}

var results map[string]interface{}

func (m mockApi) Run() error {
	results = make(map[string]interface{})
	results["TheCommand"] = m.cmd
	{{- range .UserSettings}}
	{{- if or (ne .Cmdline "") (ne .CmdlineShortcut "")}}
	results["{{.NameCode}}"] = m.config.{{.NameCode}}()
	{{- end}}
	{{- end}}
	return nil
}

func registerMockApi(c container.Container) {
	{{- range .Commands}}
	c.Singleton(func(config settings.Settings) api.{{.ApiName}}Api {return mockApi{"{{.UseName}}",config}})
	{{- end}}
}


