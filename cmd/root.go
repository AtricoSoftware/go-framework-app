package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "alternate config file")
}

func initConfig() {
	// Config file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := tryReadConfig(); err != nil {
			// Fail if specified config cannot be read
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Standard name for config
		viper.SetConfigName(fmt.Sprintf(".%s", pkg.Name))
		// Try current working directory
		dir, err := os.Getwd()
		if err == nil {
			viper.AddConfigPath(dir)
			err = tryReadConfig()
		}
		if err != nil {
			// Finally, try home directory
			dir, err = homedir.Dir()
			if err == nil {
				viper.AddConfigPath(dir)
				tryReadConfig()
			}
		}
	}
}

func tryReadConfig() error {
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		todo := viper.AllSettings()
		println(todo)
	}
	return err
}
