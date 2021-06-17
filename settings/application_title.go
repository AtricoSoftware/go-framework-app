// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationTitleSettingName = "Application.Title"
const applicationTitleSettingCmdline = "title"
const applicationTitleSettingShortcut = 't'

// Cached value
var applicationTitleSettingCache = NewCachedStringValue(func() string { return viper.GetString(applicationTitleSettingName) })

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return applicationTitleSettingCache.GetValue()
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationTitleSettingName, "Name of application").Cmdline(applicationTitleSettingCmdline).CmdlineShortcut(applicationTitleSettingShortcut).AddTo(flagSet)
}

// SECTION-END
