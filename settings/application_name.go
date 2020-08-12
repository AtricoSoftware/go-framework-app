package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationNameSettingName = "name"
const applicationNameSettingShortcut = "n"

// Fetch the setting
func (theSettings) ApplicationName() string {
	return viper.GetString(applicationNameSettingName)
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationNameSettingName, applicationNameSettingShortcut, "Name of application")
}