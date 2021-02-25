// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"


// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	if !userSettingsSettingLazy.hasValue {
		setting := viper.Get(userSettingsSettingName)
		userSettingsSettingLazy.theValue = ParseUserSettingsSetting(setting)
		userSettingsSettingLazy.hasValue = true
	}
	return userSettingsSettingLazy.theValue
}
// Lazy value
var userSettingsSettingLazy struct {
	theValue []UserSetting
	hasValue bool
}
