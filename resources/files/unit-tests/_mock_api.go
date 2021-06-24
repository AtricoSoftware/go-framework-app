// {{.Comment}}
package unit_tests

import (
	"github.com/atrico-go/container"

	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/settings"
)

type mockApi struct {
	cmd    []string
	args   []string
	config settings.Settings
}
type mockApiFactory mockApi

var results map[string]interface{}

func (m mockApi) Run() error {
	results = make(map[string]interface{})
	results["TheCommand"] = m.cmd
	results["Args"] = m.args
	{{- range .UserSettings}}
	{{- if or (ne .Cmdline "") (ne .CmdlineShortcut "")}}
	results["{{.NameCode}}"] = m.config.{{.NameCode}}()
	{{- end}}
	{{- end}}
	return nil
}

func (f mockApiFactory) Create(args []string) api.Runnable {
	f.args = args
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
{{- range .Commands}}
{{- if not .NoImplementation}}
	c.Singleton(func(config settings.Settings) api.{{.ApiName}}ApiFactory { return mockApiFactory{[]string{ {{- commaList (quoted .SplitPath) -}} }, nil, config} })
{{- end}}
{{- end}}
}