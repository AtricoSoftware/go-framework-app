// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationDescriptionSettingName = "Application.Description"
const applicationDescriptionSettingCmdline = "description"

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	return viper.GetString(applicationDescriptionSettingName)
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationDescriptionSettingName, "Description of application").Cmdline(applicationDescriptionSettingCmdline).AddTo(flagSet)
}
// SECTION-END
