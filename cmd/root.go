package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/pkg"
)

var rootCmd = &cobra.Command{
	Use:   pkg.Name,
	Short: pkg.Summary,
	Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
}

var fullVersion bool

var showVersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Shows version",
	Run: func(*cobra.Command, []string) {
		if fullVersion {
			fmt.Println(pkg.Name)
			fmt.Println(pkg.Description)
		}
		fmt.Println(pkg.Version)
		if fullVersion {
			fmt.Printf("Git: %s\n", pkg.Git)
			fmt.Printf("Built on: %s\n", pkg.BuiltOn)
			fmt.Printf("Built by: %s\n", pkg.BuiltBy)
		}
	},
}

func init() {
	showVersionCommand.PersistentFlags().BoolVarP(&fullVersion, "full", "f", false, "Full program information")
	rootCmd.AddCommand(showVersionCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
