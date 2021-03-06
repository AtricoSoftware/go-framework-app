// {{.Comment}}
package main

import (
	"fmt"
	"os"

	"github.com/atrico-go/container"
	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/cmd"
	"{{.RepositoryPath}}/settings"
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
	settings.RegisterVerboseService(c)
	cmd.RegisterCmd(c)
	api.RegisterApiFactories(c)
	return c
}
