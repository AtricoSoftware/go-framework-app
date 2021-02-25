// Generated 2021-02-25 11:57:44 by go-framework development-version
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
var repositoryPathSettingLazy struct {
	theValue string
	hasValue bool
}

// Fetch the setting
func (theSettings) RepositoryPath() string {
	if !repositoryPathSettingLazy.hasValue {
		repositoryPathSettingLazy.theValue = viper.GetString(repositoryPathSettingName)
		repositoryPathSettingLazy.hasValue = true
	}
	return repositoryPathSettingLazy.theValue
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingCmdline, repositoryPathSettingShortcut, "Path to repository")
}
