// Generated 2021-02-24 17:16:41 by go-framework development-version
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
	return viper.GetString(applicationSummarySettingName)
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationSummarySettingName, applicationSummarySettingCmdline, "Summary description of application")
}
