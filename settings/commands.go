// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"


// Fetch the setting
func (theSettings) Commands() []UserCommand {
	if !commandsSettingLazy.hasValue {
		setting := viper.Get(commandsSettingName)
		commandsSettingLazy.theValue = ParseCommandsSetting(setting)
		commandsSettingLazy.hasValue = true
	}
	return commandsSettingLazy.theValue
}
// Lazy value
var commandsSettingLazy struct {
	theValue []UserCommand
	hasValue bool
}
