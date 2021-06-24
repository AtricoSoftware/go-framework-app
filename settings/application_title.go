// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
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
