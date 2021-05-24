// {{.Comment}}
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"
)

// Create all the commands
type CommandFactory interface {
	Create() *cobra.Command
}

type commandFactory []*cobra.Command

// Register Commands
func RegisterCmd(c container.Container) {
{{- range .Commands}}
	RegisterCmd{{.ApiName}}(c)
{{- end}}
	c.Singleton(func({{- range .Commands}}{{.LowerApiName}} {{.ApiName}}Cmd, {{- end}}) CommandFactory {
		return commandFactory{
{{- range.Commands}}
			{{.LowerApiName}},
{{- end}}
		}
	})
}

func (c commandFactory) Create() *cobra.Command {
	cobra.OnInitialize(initConfig)
	rootCmd := createRootCommand()
	rootCmd.AddCommand(createVersionCommand())
	for _,cmd := range c {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}
