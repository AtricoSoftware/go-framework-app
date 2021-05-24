// Generated 2021-03-30 15:32:41 by go-framework development-version
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type GenerateCmd commandInfo

func RegisterCmdGenerate(c container.Container) {
	c.Singleton(func(apiFactory api.GenerateApiFactory) GenerateCmd { return GenerateCmd(createGenerateCommand(apiFactory)) })
}

func createGenerateCommand(apiFactory api.Factory) commandInfo {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate framework app",
		RunE: func(*cobra.Command, []string) error {
			theApi := apiFactory.Create()
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
