package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/atrico-go/viperEx"
)

// This is the name by which the setting is specified on the commandline
const applicationDescriptionSettingName = "description"
const applicationDescriptionSettingDefault = "TODO"

// Fetch the setting
func (theSettings) ApplicationDescription() string {
	return viper.GetString(applicationDescriptionSettingName)
}

func AddApplicationDescriptionFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSettingD(flagSet, applicationDescriptionSettingName, applicationDescriptionSettingDefault, "Description of application")
}
