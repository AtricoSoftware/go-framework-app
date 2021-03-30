// Generated 2021-03-30 15:32:41 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const repositoryPathSettingName = "Application.Repository"
const repositoryPathSettingCmdline = "repository"
const repositoryPathSettingShortcut = "r"

// Cached value
var repositoryPathSettingCache = NewCachedStringValue(func() string { return viper.GetString(repositoryPathSettingName) })

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return repositoryPathSettingCache.GetValue()
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingCmdline, repositoryPathSettingShortcut, "Path to repository")
}
// SECTION-END
