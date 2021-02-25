// Generated 2021-02-25 15:41:38 by go-framework development-version
package settings


import (
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"

// Lazy value
var userSettingsSettingLazy = NewLazyUserSettingSliceValue(func () []UserSetting { return ParseUserSettingsSetting(viper.Get(userSettingsSettingName)) })

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	return userSettingsSettingLazy.GetValue()
}
