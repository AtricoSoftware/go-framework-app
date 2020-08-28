package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/atrico-go/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationSummarySettingName = "summary"
const applicationSummarySettingDefault = "TODO"

// Fetch the setting
func (theSettings) ApplicationSummary() string {
	return viper.GetString(applicationSummarySettingName)
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingD(flagSet, applicationSummarySettingName, applicationSummarySettingDefault, "Summary of application")
}
