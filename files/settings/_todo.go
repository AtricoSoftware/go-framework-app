package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"{{.RepositoryPath}}/viperEx"
)

// This is the name by which the setting is specified on the commandline
const exampleSettingName = "example-setting"

// Fetch the setting
func (theSettings) Example() string {
	return viper.GetString(exampleSettingName)
}

func AddExampleFlag(flagSet *pflag.FlagSet) {
	viperEx.AddStringSetting(flagSet, exampleSettingName, "Description of setting")
	// Use P version if you want a shorthand option
	//	viperEx.AddStringSettingP(flagSet, exampleSettingName, "e", "Description of setting")
}
