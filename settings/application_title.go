// Generated 2021-02-25 15:41:38 by go-framework development-version
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
var applicationTitleSettingLazy = NewLazyStringValue(func () string { return viper.GetString(applicationTitleSettingName) })

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return applicationTitleSettingLazy.GetValue()
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingCmdline, applicationTitleSettingShortcut, "Name of application")
}
