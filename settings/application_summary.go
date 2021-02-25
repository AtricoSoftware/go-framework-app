// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationSummarySettingName = "Application.Summary"
const applicationSummarySettingCmdline = "summary"


// Fetch the setting
func (theSettings) ApplicationSummary() string {
	if !applicationSummarySettingLazy.hasValue {
		applicationSummarySettingLazy.theValue = viper.GetString(applicationSummarySettingName)
		applicationSummarySettingLazy.hasValue = true
	}
	return applicationSummarySettingLazy.theValue
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationSummarySettingName, applicationSummarySettingCmdline, "Summary description of application")
}
// Lazy value
var applicationSummarySettingLazy struct {
	theValue string
	hasValue bool
}
