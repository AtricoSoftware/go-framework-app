// Generated 2021-02-24 17:16:41 by go-framework development-version
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