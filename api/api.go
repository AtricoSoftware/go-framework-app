// Generated 2021-06-04 15:53:11 by go-framework development-version
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
