// Generated 2021-02-25 11:49:25 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationNameSettingName = "Application.Name"
const applicationNameSettingCmdline = "name"
const applicationNameSettingShortcut = "n"

// Fetch the setting
func (theSettings) ApplicationName() string {
	return viper.GetString(applicationNameSettingName)
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationNameSettingName, applicationNameSettingCmdline, applicationNameSettingShortcut, "Name of application")
}
