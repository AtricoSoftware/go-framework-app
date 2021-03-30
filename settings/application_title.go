// Generated 2021-03-30 15:32:41 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationTitleSettingName = "Application.Title"
const applicationTitleSettingCmdline = "title"
const applicationTitleSettingShortcut = "t"

// Cached value
var applicationTitleSettingCache = NewCachedStringValue(func() string { return viper.GetString(applicationTitleSettingName) })

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return applicationTitleSettingCache.GetValue()
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingCmdline, applicationTitleSettingShortcut, "Name of application")
}
// SECTION-END
