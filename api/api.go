// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
package api

import (
	"github.com/atrico-go/container"
)

// Api command to run
type ApiCommand interface {
	Run() error
}

// Generate framework app
type GenerateApi ApiCommand

// Register Api services
func RegisterApi(c container.Container) {
	RegisterGenerate(c)
}
