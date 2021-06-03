// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationNameSettingName = "Application.Name"
const applicationNameSettingCmdline = "name"
const applicationNameSettingShortcut = 'n'

// Cached value
var applicationNameSettingCache = NewCachedStringValue(func() string { return viper.GetString(applicationNameSettingName) })

// Fetch the setting
func (theSettings) ApplicationName() string {
	return applicationNameSettingCache.GetValue()
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationNameSettingName, "Name of application").Cmdline(applicationNameSettingCmdline).CmdlineShortcut(applicationNameSettingShortcut).AddTo(flagSet)
}

// SECTION-END
