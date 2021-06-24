// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"
)

type ListCmd commandInfo

func RegisterCmdList(c container.Container) {
	c.Singleton(func() ListCmd { return ListCmd(createListCommand()) })
}

func createListCommand() commandInfo {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List values",
	}
	return commandInfo{cmd, "list"}
}
