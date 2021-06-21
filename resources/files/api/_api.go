// {{.Comment}}
package api

import (
	"github.com/atrico-go/container"
)
{{- if .IncludeDryRun}}

var DryRun bool
{{- end }}

type Runnable interface {
	Run() error
}

type Factory interface {
	Create(args []string) Runnable
}

func RegisterApiFactories(c container.Container) {
{{- range .Commands}}
{{- if not .NoImplementation}}
	RegisterVerboseService(c)
	c.Singleton(func() {{.ApiName}}ApiFactory { return {{.LowerApiName}}ApiFactory{c} })
{{- end}}
{{- end}}
}

