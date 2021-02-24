// Generated 2021-02-24 20:54:12 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"

// Lazy value
var commandsSettingLazy struct {
	theValue []UserCommand
	hasValue bool
}

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	if !commandsSettingLazy.hasValue {
		setting := viper.Get(commandsSettingName)
		commandsSettingLazy.theValue = ParseCommandsSetting(setting)
		commandsSettingLazy.hasValue = true
	}
	return commandsSettingLazy.theValue
}
