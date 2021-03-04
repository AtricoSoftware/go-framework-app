// {{.Comment}}
package api

import (
	"github.com/atrico-go/container"
)

// Api command to run
type ApiCommand interface {
	Run() error
}
{{- range .Commands}}

// {{.Description}}
type {{.ApiName}}Api ApiCommand
{{- end}}

// Register Api services
func RegisterApi(c container.Container) {
{{- range .Commands}}
	Register{{.ApiName}}(c)
{{- end}}
}
