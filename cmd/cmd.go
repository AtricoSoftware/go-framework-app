// Generated 2021-03-09 17:48:01 by go-framework development-version
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
	RegisterCmdGenerate(c)
	c.Singleton(func(generate GenerateCmd) CommandFactory {
		return commandFactory{
			generate,
		}
	})
}

func (c commandFactory) Create() *cobra.Command {
	cobra.OnInitialize(initConfig)
	rootCmd := createRootCommand()
	rootCmd.AddCommand(createVersionCommand())
	for _, cmd := range c {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}
