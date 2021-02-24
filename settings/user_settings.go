package settings

import "github.com/spf13/viper"

const userSettingsSettingName = "settings"

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	setting := viper.Get(userSettingsSettingName)
	return ParseUserSettingsSetting(setting)
}
