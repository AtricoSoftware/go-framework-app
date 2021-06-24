// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type ListSkeletonsCmd commandInfo

func RegisterCmdListSkeletons(c container.Container) {
	c.Singleton(func(apiFactory api.ListSkeletonsApiFactory) ListSkeletonsCmd {
		return ListSkeletonsCmd(createListSkeletonsCommand(apiFactory))
	})
}

func createListSkeletonsCommand(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "skeletons",
		Short: "List the available settings skeletons",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			theApi := apiFactory.Create(args)
			return theApi.Run()
		},
	}
	settings.AddSkeletonFilesFlag(cmd.PersistentFlags())
	return commandInfo{cmd, "list/skeletons"}
}
