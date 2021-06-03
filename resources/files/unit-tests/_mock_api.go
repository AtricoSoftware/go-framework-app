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
type mockApiFactory mockApi

var results map[string]interface{}

func (m mockApi) Run(args []string) error {
	results = make(map[string]interface{})
	results["TheCommand"] = m.cmd
	{{- range .UserSettings}}
	{{- if or (ne .Cmdline "") (ne .CmdlineShortcut "")}}
	results["{{.NameCode}}"] = m.config.{{.NameCode}}()
	{{- end}}
	{{- end}}
	results["Args"] = args
	return nil
}

func (f mockApiFactory) Create() api.Runnable {
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
	{{- range .Commands}}
	c.Singleton(func(config settings.Settings) api.{{.ApiName}}ApiFactory {return mockApiFactory{"{{.UseName}}",config}})
	{{- end}}
}


