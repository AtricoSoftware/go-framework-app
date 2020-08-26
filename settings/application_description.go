package settings

import (
	"github.com/spf13/pflag"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationDescriptionSettingName = "description"
const applicationDescriptionSettingDefault = "TODO"

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	return viperEx.GetStringOrDefault(applicationDescriptionSettingName, applicationDescriptionSettingDefault)
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationDescriptionSettingName, "Description of application")
}
