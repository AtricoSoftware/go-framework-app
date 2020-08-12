package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/settings"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "summary",
	Run: func(*cobra.Command, []string) {
		// Implementation here!
		settings := settings.GetSettings() // Get the default settings
		fmt.Printf("Backup = %s\n", settings.Example())
	},
}

func init() {
	settings.AddExampleFlag(exampleCmd.PersistentFlags())
	rootCmd.AddCommand(exampleCmd)
}
