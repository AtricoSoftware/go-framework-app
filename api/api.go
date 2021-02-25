// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
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
