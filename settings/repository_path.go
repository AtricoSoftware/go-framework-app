package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

// This is the name by which the setting is specified on the commandline
const repositoryPathSettingName = "repository"
const repositoryPathSettingShortcut = "r"

// Fetch the setting
func (theSettings) RepositoryPath() string {
	return viper.GetString(repositoryPathSettingName)
}

func AddRepositoryPathFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, repositoryPathSettingName, repositoryPathSettingShortcut, "Path to repository")
}
