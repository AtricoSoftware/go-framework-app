// Generated 2021-02-24 20:54:12 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"

// Lazy value
var userSettingsSettingLazy struct {
	theValue []UserSetting
	hasValue bool
}

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	if !userSettingsSettingLazy.hasValue {
		setting := viper.Get(userSettingsSettingName)
		userSettingsSettingLazy.theValue = ParseUserSettingsSetting(setting)
		userSettingsSettingLazy.hasValue = true
	}
	return userSettingsSettingLazy.theValue
}
