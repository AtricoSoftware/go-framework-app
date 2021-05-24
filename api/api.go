// Generated 2021-05-24 17:41:23 by go-framework development-version
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
