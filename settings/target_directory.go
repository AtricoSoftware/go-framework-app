// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const targetDirectorySettingName = "Config.TargetDirectory"
const targetDirectorySettingCmdline = "directory"
const targetDirectorySettingShortcut = 'd'
const targetDirectorySettingDefaultVal = "."

// Cached value
var targetDirectorySettingCache = NewCachedStringValue(func() string { return viper.GetString(targetDirectorySettingName) })

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return targetDirectorySettingCache.GetValue()
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(targetDirectorySettingName, "Target directory").Cmdline(targetDirectorySettingCmdline).CmdlineShortcut(targetDirectorySettingShortcut).DefaultVal(targetDirectorySettingDefaultVal).AddTo(flagSet)
}

// SECTION-END
