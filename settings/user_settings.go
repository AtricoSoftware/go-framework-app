// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
package settings

import (
	"github.com/spf13/viper"
)

const userSettingsSettingName = "UserSettings"

// Lazy value
var userSettingsSettingLazy = NewLazyUserSettingSliceValue(func() []UserSetting { return ParseUserSettingsSetting(viper.Get(userSettingsSettingName)) })

// Fetch the setting
func (theSettings) UserSettings() []UserSetting {
	return userSettingsSettingLazy.GetValue()
}
