// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package unit_tests

import (
	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
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
	results["TargetDirectory"] = m.config.TargetDirectory()
	results["ApplicationTitle"] = m.config.ApplicationTitle()
	results["ApplicationName"] = m.config.ApplicationName()
	results["ApplicationSummary"] = m.config.ApplicationSummary()
	results["ApplicationDescription"] = m.config.ApplicationDescription()
	results["RepositoryPath"] = m.config.RepositoryPath()
	return nil
}

func (f mockApiFactory) Create(args []string) api.Runnable {
	f.args = args
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
	c.Singleton(func(config settings.Settings) api.GenerateApiFactory { return mockApiFactory{[]string{"generate"}, nil, config} })
}
