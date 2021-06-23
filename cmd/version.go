// Generated 2021-06-23 15:07:34 by go-framework v1.21.0
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/settings"
	"github.com/spf13/cobra"
)

func createVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Shows version",
		Run: func(*cobra.Command, []string) {
			showVersion()
		},
	}
	return cmd
}

func showVersion() {
	verboseService := settings.GetVerboseService()
	verboseService.VerbosePrintln(pkg.Name)
	verboseService.VerbosePrintln(pkg.Description)
	fmt.Println(pkg.Version)
	verboseService.VerbosePrintln()
	var details map[string]interface{}
	if err := json.Unmarshal([]byte(pkg.BuildDetails), &details); err == nil && len(details) > 0 {
		verboseService.VerbosePrintln("Details")
		verboseService.VerbosePrintln("-------")
		displaySection(details, "", verboseService)
	}
}

func displaySection(section map[string]interface{}, indent string, verboseService settings.VerboseService) {
	for k, v := range section {
		verboseService.VerbosePrintf("%s%s:", indent, k)
		switch v.(type) {
		case map[string]interface{}:
			verboseService.VerbosePrintln()
			displaySection(v.(map[string]interface{}), indent+"  ", verboseService)
		default:
			verboseService.VerbosePrintln(" ", v)
		}
	}
}
