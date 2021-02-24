package settings

import (
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	setting := viper.Get(commandsSettingName)
	return ParseCommandsSetting(setting)
}

