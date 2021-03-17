// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/go-framework-app/pkg"
)

var fullVersion bool

func createVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Shows version",
		Run: func(*cobra.Command, []string) {
			showVersion(fullVersion)
		},
	}
	cmd.PersistentFlags().BoolVarP(&fullVersion, "full", "f", false, "Full program information")
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
