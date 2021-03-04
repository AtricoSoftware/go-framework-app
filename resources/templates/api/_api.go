{"Type":"Mixed", "Name":"%s"}
// {{.Comment}}
// SECTION-START: Framework
package api

import (
	"github.com/atrico-go/container"

  	"{{.RepositoryPath}}/settings"
)
// SECTION-END

func Register{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings) {{.Command.ApiName}}Api {return {{.Command.Name}}Api{config: config}})
}

type {{.Command.Name}}Api struct {
	config settings.Settings
}

// {{.Command.Description}}
func (svc {{.Command.Name}}Api) Run() error {
	// Implementation here!
	return nil
}
