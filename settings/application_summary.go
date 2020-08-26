package settings

import (
	"github.com/spf13/pflag"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationSummarySettingName = "summary"
const applicationSummarySettingDefault = "TODO"

// Fetch the setting
func (theSettings) ApplicationSummary() string {
	return viperEx.GetStringOrDefault(applicationSummarySettingName, applicationSummarySettingDefault)
}

func AddApplicationSummaryFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationSummarySettingName, "Summary of application")
}
