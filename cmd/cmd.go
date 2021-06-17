// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package cmd

import (
	"path"

	"github.com/atrico-go/container"
	"github.com/spf13/cobra"
)

// Create all the commands
type CommandFactory interface {
	Create() *cobra.Command
}

type commandInfo struct {
	cmd  *cobra.Command
	path string
}
type commandFactory []commandInfo

// Register Commands
func RegisterCmd(c container.Container) {
	RegisterCmdGenerate(c)
	c.Singleton(func(generate GenerateCmd, ) CommandFactory {
		return commandFactory{
			commandInfo(generate),
		}
	})
}

func (c commandFactory) Create() *cobra.Command {
	cobra.OnInitialize(initConfig)
	commands := make(map[string]*cobra.Command, 1)
	commands["."] = createRootCommand()
	commands["."].AddCommand(createVersionCommand())
	for _, cmdInfo := range c {
		parent := path.Dir(cmdInfo.path)
		commands[parent].AddCommand(cmdInfo.cmd)
		commands[cmdInfo.path] = cmdInfo.cmd
	}
	return commands["."]
}
