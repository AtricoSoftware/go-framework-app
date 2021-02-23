package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/atrico-go/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationTitleSettingName = "title"
const applicationTitleSettingShortcut = "t"

// Fetch the setting
func (theSettings) ApplicationTitle() string {
	return viper.GetString(applicationTitleSettingName)
}

func AddApplicationTitleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingP(flagSet, applicationTitleSettingName, applicationTitleSettingShortcut, "Title of application")
}
