// Generated 2021-02-24 17:16:41 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const librariesSettingName = "Libraries"

// Fetch the setting
func (theSettings) Libraries() map[string]string {
	setting := viper.Get(librariesSettingName)
	return ParseLibrariesSetting(setting)
}
