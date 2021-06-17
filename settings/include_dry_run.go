// Generated 2021-06-17 17:07:26 by go-framework v1.20.0
// SECTION-START: Framework
package settings

import (
	"github.com/spf13/viper"
)

const includeDryRunSettingName = "Config.IncludeDryRun"
const includeDryRunSettingDefaultVal = true

// Cached value
var includeDryRunSettingCache = NewCachedBoolValue(func() bool { return viper.GetBool(includeDryRunSettingName) })

// Fetch the setting
func (theSettings) IncludeDryRun() bool {
	return includeDryRunSettingCache.GetValue()
}

func init() {
	viper.SetDefault(includeDryRunSettingName, includeDryRunSettingDefaultVal)
}

// SECTION-END
