// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationDescriptionSettingName = "Application.Description"
const applicationDescriptionSettingCmdline = "description"

// Cached value
var applicationDescriptionSettingCache = NewCachedStringValue(func() string { return viper.GetString(applicationDescriptionSettingName) })

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	return applicationDescriptionSettingCache.GetValue()
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, applicationDescriptionSettingName, applicationDescriptionSettingCmdline, "Description of application")
}

// SECTION-END
