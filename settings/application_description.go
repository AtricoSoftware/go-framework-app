// Generated 2021-02-25 11:57:44 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationDescriptionSettingName = "Application.Description"
const applicationDescriptionSettingCmdline = "description"

// Lazy value
var applicationDescriptionSettingLazy struct {
	theValue string
	hasValue bool
}

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	if !applicationDescriptionSettingLazy.hasValue {
		applicationDescriptionSettingLazy.theValue = viper.GetString(applicationDescriptionSettingName)
		applicationDescriptionSettingLazy.hasValue = true
	}
	return applicationDescriptionSettingLazy.theValue
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationDescriptionSettingName, applicationDescriptionSettingCmdline, "Description of application")
}
