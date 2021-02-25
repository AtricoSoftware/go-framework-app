// Generated 2021-02-25 15:41:38 by go-framework development-version
package settings


import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const targetDirectorySettingName = "Config.TargetDirectory"
const targetDirectorySettingCmdline = "directory"
const targetDirectorySettingShortcut = "d"
const targetDirectorySettingDefaultVal = "."

// Lazy value
var targetDirectorySettingLazy = NewLazyStringValue(func () string { return viper.GetString(targetDirectorySettingName) })

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return targetDirectorySettingLazy.GetValue()
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, targetDirectorySettingName, targetDirectorySettingCmdline, targetDirectorySettingShortcut, "Target directory")
}

func init() {
	viper.SetDefault(targetDirectorySettingName, targetDirectorySettingDefaultVal)
}
