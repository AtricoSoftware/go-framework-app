// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationTitleSettingName = "Application.Title"
const applicationTitleSettingCmdline = "title"
const applicationTitleSettingShortcut = 't'

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return viper.GetString(applicationTitleSettingName)
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationTitleSettingName, "Name of application").Cmdline(applicationTitleSettingCmdline).CmdlineShortcut(applicationTitleSettingShortcut).AddTo(flagSet)
}
// SECTION-END
