// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

func CreateGenerateCommand(c container.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate framework app",
		RunE: func(*cobra.Command, []string) error {
			var generateApi api.GenerateApi
			c.Make(&generateApi)
			return generateApi.Run()
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
