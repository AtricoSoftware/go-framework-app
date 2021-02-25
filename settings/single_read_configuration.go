// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
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
