// Generated 2021-03-30 15:32:41 by go-framework development-version
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/pkg"
)

func createVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Shows version",
		Run: func(*cobra.Command, []string) {
			showVersion(api.VerboseFlag)
		},
	}
	return cmd
}

func showVersion(fullVersion bool) {
	if fullVersion {
		fmt.Println(pkg.Name)
		fmt.Println(pkg.Description)
	}
	fmt.Println(pkg.Version)
	if fullVersion {
		fmt.Println()
		var details map[string]interface{}
		if err := json.Unmarshal([]byte(pkg.BuildDetails), &details); err == nil && len(details) > 0 {
			fmt.Println("Details")
			fmt.Println("-------")
			displaySection(details, "")
		}
	}
}

func displaySection(section map[string]interface{}, indent string) {
	for k, v := range section {
		fmt.Printf("%s%s:", indent, k)
		switch v.(type) {
		case map[string]interface{}:
			fmt.Println()
			displaySection(v.(map[string]interface{}), indent+"  ")
		default:
			fmt.Printf(" %s\n", v)
		}
	}
}
