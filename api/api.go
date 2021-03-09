// Generated 2021-03-09 17:48:01 by go-framework development-version
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run() error
}

// Register Api services
func RegisterApi(c container.Container) {
	RegisterApiGenerate(c)
}
