// {{.Comment}}
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"{{.RepositoryPath}}/pkg"
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
	verbosePrintln(pkg.Name)
	verbosePrintln(pkg.Description)
	fmt.Println(pkg.Version)
	verbosePrintln()
	var details map[string]interface{}
	if err := json.Unmarshal([]byte(pkg.BuildDetails), &details); err == nil && len(details) > 0 {
		verbosePrintln("Details")
		verbosePrintln("-------")
		displaySection(details, "")
	}
}

func displaySection(section map[string]interface{}, indent string) {
	for k, v := range section {
		verbosePrintf("%s%s:", indent, k)
		switch v.(type) {
		case map[string]interface{}:
			verbosePrintln()
			displaySection(v.(map[string]interface{}), indent+"  ")
		default:
			verbosePrintln(" ",v)
		}
	}
}
