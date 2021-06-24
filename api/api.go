// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
package api

import (
	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/settings"
)

type Runnable interface {
	Run() error
}

type Factory interface {
	Create(args []string) Runnable
}

func RegisterApiFactories(c container.Container) {
	settings.RegisterVerboseService(c)
	c.Singleton(func() GenerateApiFactory { return generateApiFactory{c} })
	c.Singleton(func() ListSkeletonsApiFactory { return listSkeletonsApiFactory{c} })
}
