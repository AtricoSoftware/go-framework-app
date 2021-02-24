// Generated 2021-02-24 17:16:41 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const repositoryPathSettingName = "Application.Repository"
const repositoryPathSettingCmdline = "repository"
const repositoryPathSettingShortcut = "r"

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return viper.GetString(repositoryPathSettingName)
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingCmdline, repositoryPathSettingShortcut, "Path to repository")
}
