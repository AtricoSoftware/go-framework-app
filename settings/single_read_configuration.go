// Generated 2021-03-30 15:32:41 by go-framework development-version
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
