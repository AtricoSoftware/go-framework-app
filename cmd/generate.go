// Generated 2021-03-09 17:48:01 by go-framework development-version
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

type GenerateCmd *cobra.Command

func RegisterCmdGenerate(c container.Container) {
	c.Singleton(func(api api.GenerateApi) GenerateCmd { return createGenerateCommand(api) })
}

func createGenerateCommand(api api.Runnable) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate framework app",
		RunE: func(*cobra.Command, []string) error {
			return api.Run()
		},
	}
	settings.AddTargetDirectoryFlag(cmd.PersistentFlags())
	settings.AddApplicationTitleFlag(cmd.PersistentFlags())
	settings.AddApplicationNameFlag(cmd.PersistentFlags())
	settings.AddApplicationSummaryFlag(cmd.PersistentFlags())
	settings.AddApplicationDescriptionFlag(cmd.PersistentFlags())
	settings.AddRepositoryPathFlag(cmd.PersistentFlags())
	return cmd
}
