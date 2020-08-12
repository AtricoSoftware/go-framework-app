package main

//go:generate go run create_resources.go

// +build ignore

import (
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/cmd"
)

func main() {
	cmd.Execute()
}
