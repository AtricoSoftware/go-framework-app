// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
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

// Cached value
var repositoryPathSettingCache = NewCachedStringValue(func() string { return viper.GetString(repositoryPathSettingName) })

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return repositoryPathSettingCache.GetValue()
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(repositoryPathSettingName, "Path to repository").Cmdline(repositoryPathSettingCmdline).CmdlineShortcut(repositoryPathSettingShortcut).AddTo(flagSet)
}

// SECTION-END
