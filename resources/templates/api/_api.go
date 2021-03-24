{"Type":"Mixed", "Name":"%s"}
// {{.Comment}}
// SECTION-START: Framework
package api

import (
	"github.com/atrico-go/container"

  	"{{.RepositoryPath}}/settings"
)

type {{.Command.ApiName}}Api Runnable
type {{.Command.ApiName}}ApiFactory RunnableFactory
// SECTION-END

func RegisterApi{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings) {{.Command.ApiName}}ApiFactory {return {{.Command.Name}}ApiFactory{config}})
}

type {{.Command.Name}}Api struct {
	config settings.Settings
}
type {{.Command.Name}}ApiFactory {{.Command.Name}}Api

func (f {{.Command.Name}}ApiFactory) Create() Runnable {
	return {{.Command.Name}}Api(f)
}

// {{.Command.Description}}
func (svc {{.Command.Name}}Api) Run() error {
	// Implementation here!
	return nil
}
