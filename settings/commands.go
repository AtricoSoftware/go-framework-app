// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
package settings

import (
	"github.com/spf13/viper"
)

const commandsSettingName = "Commands"

// Lazy value
var commandsSettingLazy = NewLazyUserCommandSliceValue(func() []UserCommand { return ParseCommandsSetting(viper.Get(commandsSettingName)) })

// Fetch the setting
func (theSettings) Commands() []UserCommand {
	return commandsSettingLazy.GetValue()
}
