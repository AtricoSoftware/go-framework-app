// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
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
