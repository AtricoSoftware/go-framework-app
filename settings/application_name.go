// Generated 2021-02-25 15:41:38 by go-framework development-version
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
var applicationNameSettingLazy = NewLazyStringValue(func () string { return viper.GetString(applicationNameSettingName) })

// Fetch the setting
func (theSettings) ApplicationName() string {
	return applicationNameSettingLazy.GetValue()
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationNameSettingName, applicationNameSettingCmdline, applicationNameSettingShortcut, "Name of application")
}
