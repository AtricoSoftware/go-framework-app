// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AtricoSoftware/go-framework-app/api"
	"github.com/AtricoSoftware/go-framework-app/pkg"
)

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   pkg.Name,
		Short: pkg.Summary,
		Long:  fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "alternate config file")
	cmd.PersistentFlags().BoolVarP(&api.VerboseFlag, "verbose", "", false, "More output")
	return cmd
}

var cfgFile string

func initConfig() {
	var err error
	// Config file
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
		err := viper.ReadInConfig()
		if err == nil {
			api.VerbosePrintln("Using config file:", viper.ConfigFileUsed())
		}
	}
	return err
}
