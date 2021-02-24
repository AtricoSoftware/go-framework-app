// Generated 2021-02-24 17:16:41 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	setting := viper.Get(userSettingsSettingName)
	return ParseUserSettingsSetting(setting)
}
