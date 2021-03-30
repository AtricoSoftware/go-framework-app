// Generated 2021-03-30 15:32:41 by go-framework development-version
package api

import (
	"github.com/atrico-go/container"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create() Runnable
}

func RegisterApiFactories(c container.Container) {
	c.Singleton(func() GenerateApiFactory { return generateApiFactory{c} })
}
