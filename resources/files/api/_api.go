// {{.Comment}}
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run() error
}

type RunnableFactory interface {
	Create() Runnable
}

// Register Api services
func RegisterApi(c container.Container) {
{{- range .Commands}}
	RegisterApi{{.ApiName}}(c)
{{- end}}
}
