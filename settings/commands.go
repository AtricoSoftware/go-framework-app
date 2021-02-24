// Generated 2021-02-24 17:16:41 by go-framework development-version
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
