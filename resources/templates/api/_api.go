{"Type":"Mixed", "Name":"%s"}
// {{.Comment}}
// SECTION-START: Framework
package api

import (
	"github.com/atrico-go/container"

  	"{{.RepositoryPath}}/settings"
)

type {{.Command.ApiName}}Api Runnable
type {{.Command.ApiName}}ApiFactory Factory

type {{.Command.LowerApiName}}ApiFactory struct {
	container.Container
}

func (f {{.Command.LowerApiName}}ApiFactory) Create() Runnable {
	RegisterApi{{.Command.ApiName}}(f.Container)
	var theApi {{.Command.ApiName}}Api
	f.Container.Make(&theApi)
	return theApi
}
// SECTION-END

func RegisterApi{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings) {{.Command.ApiName}}Api {return {{.Command.LowerApiName}}Api{config}})
}

type {{.Command.LowerApiName}}Api struct {
	settings.Settings
}

// {{.Command.Description}}
func (svc {{.Command.LowerApiName}}Api) Run(args []string) error {
	// Implementation here!
	return nil
}
