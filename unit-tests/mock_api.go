// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
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
	results["SkeletonFiles"] = m.config.SkeletonFiles()
	results["ConfigFile"] = m.config.ConfigFile()
	results["Verbose"] = m.config.Verbose()
	return nil
}

func (f mockApiFactory) Create(args []string) api.Runnable {
	f.args = args
	return mockApi(f)
}

func registerMockApiFactories(c container.Container) {
	c.Singleton(func(config settings.Settings) api.GenerateApiFactory { return mockApiFactory{[]string{"generate"}, nil, config} })
	c.Singleton(func(config settings.Settings) api.ListSkeletonsApiFactory { return mockApiFactory{[]string{"list", "skeletons"}, nil, config} })
}
