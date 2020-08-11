package cmd

import (
	"github.com/spf13/cobra"

	"{{.RepositoryPath}}/settings"
)

var {{.Command.Name}}Cmd = &cobra.Command{
	Use:   "{{.Command.Name}}",
	Short: "{{.Command.Description}}",
	Run: func(*cobra.Command, []string) {
		// Implementation here!
		settings := settings.GetSettings() // Get the default settings
		println(settings)
	},
}

func init() {
	rootCmd.AddCommand({{.Command.Name}}Cmd)
{{- range .UserSettings}}
	{{- if .AppliesToCmd $.Command.Name}}
	settings.Add{{.Name}}Flag({{$.Command.Name}}Cmd.PersistentFlags())
	{{- end}}
{{- end}}
}
