// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFileSettingName = "ConfigFile"
const configFileSettingCmdline = "config-file"

// Cached value
var configFileSettingCache = NewCachedStringValue(func() string { return viper.GetString(configFileSettingName) })

// Fetch the setting
func (theSettings) ConfigFile() string {
	return configFileSettingCache.GetValue()
}

func AddConfigFileFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(configFileSettingName, "Alternate config file").Cmdline(configFileSettingCmdline).AddTo(flagSet)
}

// SECTION-END
