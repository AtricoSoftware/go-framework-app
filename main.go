// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
package main

import (
	"fmt"
	"os"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/cmd"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/atrico-go/container"
)

func main() {
	c := register()
	var cmdFactory cmd.CommandFactory
	c.Make(&cmdFactory)
	cmd := cmdFactory.Create()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func register() container.Container {
	c := container.NewContainer()
	settings.RegisterSettings(c)
	cmd.RegisterCmd(c)
	api.RegisterApiFactories(c)
	return c
}
