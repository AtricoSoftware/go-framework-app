// Generated 2021-02-25 11:57:44 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const targetDirectorySettingName = "directory"
const targetDirectorySettingCmdline = "directory"
const targetDirectorySettingShortcut = "d"
const targetDirectorySettingDefaultVal = "."

// Lazy value
var targetDirectorySettingLazy struct {
	theValue string
	hasValue bool
}

// Fetch the setting
func (theSettings) TargetDirectory() string {
	if !targetDirectorySettingLazy.hasValue {
		targetDirectorySettingLazy.theValue = viper.GetString(targetDirectorySettingName)
		targetDirectorySettingLazy.hasValue = true
	}
	return targetDirectorySettingLazy.theValue
}

func AddTargetDirectoryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, targetDirectorySettingName, targetDirectorySettingCmdline, targetDirectorySettingShortcut, "Target directory")
}

func init() {
	viper.SetDefault(targetDirectorySettingName, targetDirectorySettingDefaultVal)
}
