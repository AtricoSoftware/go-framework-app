package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/pkg"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var rootCmd = &cobra.Command{
	Use:   pkg.Name,
	Short: pkg.Summary,
	Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
}
