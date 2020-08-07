package cmd

import (
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/settings"
)

var {{.CommandName}}Cmd = &cobra.Command{
	Use:   "{{.CommandName}}",
	Short: "TODO summary",
	Run: func(*cobra.Command, []string) {
		// Implementation here!
		settings := settings.GetSettings() // Get the default settings
		println(settings)
	},
}

func init() {
	rootCmd.AddCommand({{.CommandName}}Cmd)
}
