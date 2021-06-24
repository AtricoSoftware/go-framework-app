// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const verboseSettingName = "Verbose"
const verboseSettingCmdline = "verbose"
const verboseSettingShortcut = 'v'
const verboseSettingEnvVar = "VERBOSE"

// Cached value
var verboseSettingCache = NewCachedBoolValue(func() bool { return viper.GetBool(verboseSettingName) })

// Fetch the setting
func (theSettings) Verbose() bool {
	return verboseSettingCache.GetValue()
}

func AddVerboseFlag(flagSet *pflag.FlagSet) {
	viperEx.BoolSetting(verboseSettingName, "Generate more detailed output").Cmdline(verboseSettingCmdline).CmdlineShortcut(verboseSettingShortcut).EnvVar(verboseSettingEnvVar).AddTo(flagSet)
}

// SECTION-END
