// Generated 2021-06-24 14:50:11 by go-framework v1.21.1
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
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
	viperEx.StringSetting(applicationDescriptionSettingName, "Description of application").Cmdline(applicationDescriptionSettingCmdline).AddTo(flagSet)
}

// SECTION-END
