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
