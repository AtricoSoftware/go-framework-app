// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run(args []string) error
}

type Factory interface {
	Create() Runnable
}

func RegisterApiFactories(c container.Container) {
	c.Singleton(func() GenerateApiFactory { return generateApiFactory{c} })
}
