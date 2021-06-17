// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
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

func init() {
	viper.SetDefault(singleReadConfigurationSettingName, singleReadConfigurationSettingDefaultVal)
}

// SECTION-END
