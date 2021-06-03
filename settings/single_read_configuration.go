// Generated 2021-06-03 14:15:48 by go-framework v1.17.0
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/viper"
)

const singleReadConfigurationSettingName = "Config.SingleReadConfig"
const singleReadConfigurationSettingDefaultVal = true

// Cached value
var singleReadConfigurationSettingCache = NewCachedBoolValue(func() bool { return viper.GetBool(singleReadConfigurationSettingName) })

// Fetch the setting
func (theSettings) SingleReadConfiguration() bool {
	return singleReadConfigurationSettingCache.GetValue()
}

// SECTION-END
