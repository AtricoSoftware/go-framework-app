// Generated 2021-03-09 17:48:01 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationDescriptionSettingName = "Application.Description"
const applicationDescriptionSettingCmdline = "description"

// Lazy value
var applicationDescriptionSettingLazy = NewLazyStringValue(func() string { return viper.GetString(applicationDescriptionSettingName) })

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	return applicationDescriptionSettingLazy.GetValue()
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationDescriptionSettingName, applicationDescriptionSettingCmdline, "Description of application")
}

// SECTION-END
