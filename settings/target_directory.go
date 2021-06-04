// Generated 2021-06-04 15:53:11 by go-framework development-version
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

// Fetch the setting
func (theSettings) TargetDirectory() string {
	return viper.GetString(targetDirectorySettingName)
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(targetDirectorySettingName, "Target directory").Cmdline(targetDirectorySettingCmdline).CmdlineShortcut(targetDirectorySettingShortcut).DefaultVal(targetDirectorySettingDefaultVal).AddTo(flagSet)
}
// SECTION-END
