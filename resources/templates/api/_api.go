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

func (f {{.Command.LowerApiName}}ApiFactory) Create(args []string) Runnable {
	RegisterApi{{.Command.ApiName}}(f.Container)
{{- if .Command.HasArgs}}
	theArgs := make(map[string]string)
	for i, arg := range []string{ {{- commaList (quoted (concat .Command.Args .Command.OptionalArgs)) -}} } {
		if i < len(args) {
			theArgs[arg] = args[i]
		}
	}
	var addArgs settings.SetArgs
	f.Container.Make(&addArgs)
	addArgs.SetArgs(theArgs)
{{- end}}
	var theApi {{.Command.ApiName}}Api
	f.Container.Make(&theApi)
	return theApi
}

// SECTION-END

func RegisterApi{{.Command.ApiName}}(c container.Container) {
	c.Singleton(func(config settings.Settings, verboseService VerboseService) {{.Command.ApiName}}Api {return {{.Command.LowerApiName}}Api{config, verboseService}})
}

type {{.Command.LowerApiName}}Api struct {
	settings.Settings
	VerboseService
}

// {{.Command.Description}}
func (svc {{.Command.LowerApiName}}Api) Run() error {
	// Implementation here!
	return nil
}
