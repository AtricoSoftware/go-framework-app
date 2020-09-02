package api

import (
	"github.com/atrico-go/container"

  	"{{.RepositoryPath}}/settings"
)

func Register{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings) {{.Command.ApiName}}Api {return {{.Command.Name}}Api{config: config}})
}

// {{.Command.Description}}
func (svc {{.Command.Name}}Api) Run() error {
	// Implementation here!
	return nil
}

type {{.Command.Name}}Api struct {
	config settings.Settings
}
