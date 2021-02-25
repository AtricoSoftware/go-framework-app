// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const repositoryPathSettingName = "Application.Repository"
const repositoryPathSettingCmdline = "repository"
const repositoryPathSettingShortcut = "r"

// Lazy value
var repositoryPathSettingLazy = NewLazyStringValue(func() string { return viper.GetString(repositoryPathSettingName) })

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return repositoryPathSettingLazy.GetValue()
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingCmdline, repositoryPathSettingShortcut, "Path to repository")
}
