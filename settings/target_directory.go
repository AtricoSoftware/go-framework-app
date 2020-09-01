package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/atrico-go/viperEx"
)

// This is the name by which the setting is specified on the commandline
const targetDirectorySettingName = "directory"
const targetDirectorySettingShortcut = "d"
const targetDirectorySettingDefault = "."

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return viper.GetString(targetDirectorySettingName)
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingPD(flagSet, targetDirectorySettingName, targetDirectorySettingShortcut, targetDirectorySettingDefault, "Target directory")
}
