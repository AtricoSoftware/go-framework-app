// Generated 2021-06-04 15:53:11 by go-framework development-version
package unit_tests

import (
	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type mockApi struct {
	cmd    []string
	config settings.Settings
}
type mockApiFactory mockApi

var results map[string]interface{}

func (m mockApi) Run(args []string) error {
	results = make(map[string]interface{})
	results["TheCommand"] = m.cmd
	results["TargetDirectory"] = m.config.TargetDirectory()
	results["ApplicationTitle"] = m.config.ApplicationTitle()
	results["ApplicationName"] = m.config.ApplicationName()
	results["ApplicationSummary"] = m.config.ApplicationSummary()
	results["ApplicationDescription"] = m.config.ApplicationDescription()
	results["RepositoryPath"] = m.config.RepositoryPath()
	results["Args"] = args
	return nil
}

func (f mockApiFactory) Create() api.Runnable {
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
	c.Singleton(func(config settings.Settings) api.GenerateApiFactory {return mockApiFactory{[]string{"generate"} ,config}})
}


