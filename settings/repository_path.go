// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const repositoryPathSettingName = "Application.Repository"
const repositoryPathSettingCmdline = "repository"
const repositoryPathSettingShortcut = 'r'

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return viper.GetString(repositoryPathSettingName)
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(repositoryPathSettingName, "Path to repository").Cmdline(repositoryPathSettingCmdline).CmdlineShortcut(repositoryPathSettingShortcut).AddTo(flagSet)
}
// SECTION-END
