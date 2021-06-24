package main

import (
	"fmt"
	"os"

	"{{.RepositoryPath}}/api"
	"{{.RepositoryPath}}/cmd"
)

func main() {
	cmd := cmd.CreateCommands(api.GetApi())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}