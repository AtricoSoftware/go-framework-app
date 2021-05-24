// {{.Comment}}
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create() Runnable
}

func RegisterApiFactories(c container.Container) {
{{- range .Commands}}
	c.Singleton(func() {{.ApiName}}ApiFactory { return {{.LowerApiName}}ApiFactory{c} })
{{- end}}
}