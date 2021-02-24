// Generated 2021-02-24 17:16:41 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationTitleSettingName = "Application.Title"
const applicationTitleSettingCmdline = "title"
const applicationTitleSettingShortcut = "t"

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return viper.GetString(applicationTitleSettingName)
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingCmdline, applicationTitleSettingShortcut, "Name of application")
}
