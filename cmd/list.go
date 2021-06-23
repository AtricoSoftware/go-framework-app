// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
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
