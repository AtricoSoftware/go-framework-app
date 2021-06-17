// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
package cmd

import (
	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"
)

type GenerateCmd commandInfo

func RegisterCmdGenerate(c container.Container) {
	c.Singleton(func(apiFactory api.GenerateApiFactory) GenerateCmd { return GenerateCmd(createGenerateCommand(apiFactory)) })
}

func createGenerateCommand(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate framework app",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			theApi := apiFactory.Create(args)
			return theApi.Run()
		},
	}
	settings.AddTargetDirectoryFlag(cmd.PersistentFlags())
	settings.AddApplicationTitleFlag(cmd.PersistentFlags())
	settings.AddApplicationNameFlag(cmd.PersistentFlags())
	settings.AddApplicationSummaryFlag(cmd.PersistentFlags())
	settings.AddApplicationDescriptionFlag(cmd.PersistentFlags())
	settings.AddRepositoryPathFlag(cmd.PersistentFlags())
	return commandInfo{cmd, "generate"}
}
