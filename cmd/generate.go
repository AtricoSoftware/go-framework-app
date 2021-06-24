// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
package cmd

import (
	"fmt"
	"os"

	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type GenerateCmd commandInfo

func RegisterCmdGenerate(c container.Container) {
	c.Singleton(func(apiFactory api.GenerateApiFactory) GenerateCmd {
		return GenerateCmd(createGenerateCommand(apiFactory))
	})
}

func createGenerateCommand(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate framework app",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			theApi := apiFactory.Create(args)
			if err := theApi.Run(); err != nil {
				fmt.Fprint(os.Stderr, err)
			}
		},
	}
	settings.AddTargetDirectoryFlag(cmd.PersistentFlags())
	settings.AddApplicationTitleFlag(cmd.PersistentFlags())
	settings.AddApplicationNameFlag(cmd.PersistentFlags())
	settings.AddApplicationSummaryFlag(cmd.PersistentFlags())
	settings.AddApplicationDescriptionFlag(cmd.PersistentFlags())
	settings.AddRepositoryPathFlag(cmd.PersistentFlags())
	settings.AddSkeletonFilesFlag(cmd.PersistentFlags())
	return commandInfo{cmd, "generate"}
}
