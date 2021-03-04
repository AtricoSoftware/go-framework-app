// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
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
	cmd := cmd.CreateCommands(c)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func register() container.Container {
	c := container.NewContainer()
	settings.RegisterSettings(c)
	api.RegisterApi(c)
	return c
}
