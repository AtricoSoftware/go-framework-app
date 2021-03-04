// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationSummarySettingName = "Application.Summary"
const applicationSummarySettingCmdline = "summary"

// Lazy value
var applicationSummarySettingLazy = NewLazyStringValue(func() string { return viper.GetString(applicationSummarySettingName) })

// Fetch the setting
func (theSettings) ApplicationSummary() string {
	return applicationSummarySettingLazy.GetValue()
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationSummarySettingName, applicationSummarySettingCmdline, "Summary description of application")
}

// SECTION-END
