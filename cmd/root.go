// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/pkg"
	"github.com/AtricoSoftware/go-framework-app/settings"
)

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   pkg.Name,
		Short: pkg.Summary,
		Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
	}
	settings.AddConfigFileFlag(cmd.PersistentFlags())
	settings.AddVerboseFlag(cmd.PersistentFlags())
	return cmd
}

func initConfig() {
	var err error
	// Config file
	cfgFile := viper.GetString("ConfigFile")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err = tryReadConfig(func() (string, error) { return cfgFile, nil }); err != nil {
			// Fail if specified config cannot be read
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		// Standard name for config
		viper.SetConfigName(fmt.Sprintf(".%s", pkg.Name))
		// Try current working directory
		err = tryReadConfig(func() (string, error) { return os.Getwd() })
		if err != nil {
			// Try home directory
			err = tryReadConfig(func() (string, error) { return homedir.Dir() })
		}
		if err != nil {
			// Try executable directory
			err = tryReadConfig(func() (string, error) {
				_, filename, _, _ := runtime.Caller(0)
				return filepath.Dir(filename), nil
			})
		}
	}
}

func tryReadConfig(getDir func() (dir string, err error)) error {
	var err error
	var dir string
	if dir, err = getDir(); err == nil {
		viper.AddConfigPath(dir)
		err = viper.ReadInConfig()
		if err == nil {
			settings.GetVerboseService().VerbosePrintln("Using config file:", viper.ConfigFileUsed())
		}
	}
	return err
}
