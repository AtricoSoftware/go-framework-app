// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
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
var applicationTitleSettingLazy = NewLazyStringValue(func() string { return viper.GetString(applicationTitleSettingName) })

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return applicationTitleSettingLazy.GetValue()
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingCmdline, applicationTitleSettingShortcut, "Name of application")
}
