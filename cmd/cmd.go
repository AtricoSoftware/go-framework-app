// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
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
	RegisterCmdList(c)
	RegisterCmdListSkeletons(c)
	c.Singleton(func(generate GenerateCmd, list ListCmd, listSkeletons ListSkeletonsCmd, ) CommandFactory {
		return commandFactory{
			commandInfo(generate),
			commandInfo(list),
			commandInfo(listSkeletons),
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
