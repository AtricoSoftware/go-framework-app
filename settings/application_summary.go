// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
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
	viperEx.StringSetting(applicationSummarySettingName, "Summary description of application").Cmdline(applicationSummarySettingCmdline).AddTo(flagSet)
}
// SECTION-END
