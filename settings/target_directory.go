package settings

import (
	"github.com/spf13/pflag"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

// This is the name by which the setting is specified on the commandline
const targetDirectorySettingName = "directory"
const targetDirectorySettingShortcut = "d"
const targetDirectorySettingDefault = "."

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return viperEx.GetStringOrDefault(targetDirectorySettingName, targetDirectorySettingDefault)
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, targetDirectorySettingName, targetDirectorySettingShortcut, "Target directory")
}
