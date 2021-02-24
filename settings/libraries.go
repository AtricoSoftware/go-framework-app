package settings

import (
	"github.com/spf13/viper"
)

const librariesSettingName = "libraries"

// Fetch the setting
func (theSettings) Libraries() map[string]string {
	setting := viper.Get(librariesSettingName)
	return ParseLibrariesSetting(setting)
}
