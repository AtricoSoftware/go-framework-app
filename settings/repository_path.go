// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
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

// Lazy value
var repositoryPathSettingLazy = NewLazyStringValue(func() string { return viper.GetString(repositoryPathSettingName) })

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return repositoryPathSettingLazy.GetValue()
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingCmdline, repositoryPathSettingShortcut, "Path to repository")
}

// SECTION-END
