// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
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
	viperEx.AddStringSetting(flagSet, applicationSummarySettingName, applicationSummarySettingCmdline, "Summary description of application")
}

// SECTION-END
