package main

//go:generate go run create_resources.go

// +build ignore

import (
	"github.com/AtricoSoftware/go-framework-app/cmd"
)

func main() {
	cmd.Execute()
}
