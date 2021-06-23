// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
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
	RegisterVerboseService(c)
	c.Singleton(func() GenerateApiFactory { return generateApiFactory{c} })
	RegisterVerboseService(c)
	c.Singleton(func() ListSkeletonsApiFactory { return listSkeletonsApiFactory{c} })
}
