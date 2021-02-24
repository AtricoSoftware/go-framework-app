// Generated 2021-02-24 20:54:12 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const librariesSettingName = "Libraries"

// Lazy value
var librariesSettingLazy struct {
	theValue map[string]string
	hasValue bool
}

// Fetch the setting
func (theSettings) Libraries() map[string]string {
	if !librariesSettingLazy.hasValue {
		setting := viper.Get(librariesSettingName)
		librariesSettingLazy.theValue = ParseLibrariesSetting(setting)
		librariesSettingLazy.hasValue = true
	}
	return librariesSettingLazy.theValue
}
