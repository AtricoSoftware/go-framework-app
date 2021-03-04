// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/viper"
)

const singleReadConfigurationSettingName = "Config.SingleReadConfig"
const singleReadConfigurationSettingDefaultVal = true

// Lazy value
var singleReadConfigurationSettingLazy = NewLazyBoolValue(func() bool { return viper.GetBool(singleReadConfigurationSettingName) })

// Fetch the setting
func (theSettings) SingleReadConfiguration() bool {
	return singleReadConfigurationSettingLazy.GetValue()
}

func init() {
	viper.SetDefault(singleReadConfigurationSettingName, singleReadConfigurationSettingDefaultVal)
}

// SECTION-END
