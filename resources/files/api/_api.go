// {{.Comment}}
package api

import (
	"github.com/atrico-go/container"
	"{{.RepositoryPath}}/settings"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create(args []string) Runnable
}

func RegisterApiFactories(c container.Container) {
	settings.RegisterVerboseService(c)
{{- range .Commands}}
{{- if not .NoImplementation}}
	c.Singleton(func() {{.ApiName}}ApiFactory { return {{.LowerApiName}}ApiFactory{c} })
{{- end}}
{{- end}}
}

