// Generated 2021-06-04 15:53:11 by go-framework development-version
// SECTION-START: Framework
package settings

import (
	"github.com/atrico-go/viperEx/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationNameSettingName = "Application.Name"
const applicationNameSettingCmdline = "name"
const applicationNameSettingShortcut = 'n'

// Fetch the setting
func (theSettings) ApplicationName() string {
	return viper.GetString(applicationNameSettingName)
}

func AddApplicationNameFlag(flagSet *pflag.FlagSet) {
	viperEx.StringSetting(applicationNameSettingName, "Name of application").Cmdline(applicationNameSettingCmdline).CmdlineShortcut(applicationNameSettingShortcut).AddTo(flagSet)
}
// SECTION-END
