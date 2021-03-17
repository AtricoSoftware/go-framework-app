// Generated 2021-03-17 16:07:26 by go-framework V1.8.0
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
