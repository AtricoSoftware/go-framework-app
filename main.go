// Generated 2021-03-09 17:48:01 by go-framework development-version
package main

import (
	"fmt"
	"os"

	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/cmd"
	"github.com/AtricoSoftware/go-framework-app/settings"
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
	api.RegisterApi(c)
	return c
}
