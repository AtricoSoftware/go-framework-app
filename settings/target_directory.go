// Generated 2021-03-30 15:32:41 by go-framework development-version
// SECTION-START: Framework
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

// Cached value
var targetDirectorySettingCache = NewCachedStringValue(func() string { return viper.GetString(targetDirectorySettingName) })

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return targetDirectorySettingCache.GetValue()
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, targetDirectorySettingName, targetDirectorySettingCmdline, targetDirectorySettingShortcut, "Target directory")
}

func init() {
	viper.SetDefault(targetDirectorySettingName, targetDirectorySettingDefaultVal)
}
// SECTION-END
