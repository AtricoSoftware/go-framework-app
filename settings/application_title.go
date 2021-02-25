// Generated 2021-02-25 11:57:44 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationTitleSettingName = "Application.Title"
const applicationTitleSettingCmdline = "title"
const applicationTitleSettingShortcut = "t"

// Lazy value
var applicationTitleSettingLazy struct {
	theValue string
	hasValue bool
}

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	if !applicationTitleSettingLazy.hasValue {
		applicationTitleSettingLazy.theValue = viper.GetString(applicationTitleSettingName)
		applicationTitleSettingLazy.hasValue = true
	}
	return applicationTitleSettingLazy.theValue
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingCmdline, applicationTitleSettingShortcut, "Name of application")
}
