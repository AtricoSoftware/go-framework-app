// Generated 2021-02-25 11:57:44 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationNameSettingName = "Application.Name"
const applicationNameSettingCmdline = "name"
const applicationNameSettingShortcut = "n"

// Lazy value
var applicationNameSettingLazy struct {
	theValue string
	hasValue bool
}

// Fetch the setting
func (theSettings) ApplicationName() string {
	if !applicationNameSettingLazy.hasValue {
		applicationNameSettingLazy.theValue = viper.GetString(applicationNameSettingName)
		applicationNameSettingLazy.hasValue = true
	}
	return applicationNameSettingLazy.theValue
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationNameSettingName, applicationNameSettingCmdline, applicationNameSettingShortcut, "Name of application")
}
