// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create(args []string) Runnable
}

func RegisterApiFactories(c container.Container) {
	c.Singleton(func() GenerateApiFactory { return generateApiFactory{c} })
}
