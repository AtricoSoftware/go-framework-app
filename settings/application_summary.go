// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationSummarySettingName = "Application.Summary"
const applicationSummarySettingCmdline = "summary"

// Cached value
var applicationSummarySettingCache = NewCachedStringValue(func() string { return viper.GetString(applicationSummarySettingName) })

// Fetch the setting
func (theSettings) ApplicationSummary() string {
	return applicationSummarySettingCache.GetValue()
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationSummarySettingName, "Summary description of application").Cmdline(applicationSummarySettingCmdline).AddTo(flagSet)
}

// SECTION-END
