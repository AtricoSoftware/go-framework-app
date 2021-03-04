// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
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
