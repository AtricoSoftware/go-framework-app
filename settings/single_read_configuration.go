// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import (
	"github.com/spf13/viper"
)

const singleReadConfigurationSettingName = "Config.SingleReadConfig"
const singleReadConfigurationSettingDefaultVal = true


// Fetch the setting
func (theSettings) SingleReadConfiguration() bool {
	if !singleReadConfigurationSettingLazy.hasValue {
		singleReadConfigurationSettingLazy.theValue = viper.GetBool(singleReadConfigurationSettingName)
		singleReadConfigurationSettingLazy.hasValue = true
	}
	return singleReadConfigurationSettingLazy.theValue
}

func init() {
	viper.SetDefault(singleReadConfigurationSettingName, singleReadConfigurationSettingDefaultVal)
}
// Lazy value
var singleReadConfigurationSettingLazy struct {
	theValue bool
	hasValue bool
}
