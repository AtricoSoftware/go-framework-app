// Generated 2021-02-25 15:41:38 by go-framework development-version
package settings


import (
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"

// Lazy value
var commandsSettingLazy = NewLazyUserCommandSliceValue(func () []UserCommand { return ParseCommandsSetting(viper.Get(commandsSettingName)) })

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	return commandsSettingLazy.GetValue()
}
