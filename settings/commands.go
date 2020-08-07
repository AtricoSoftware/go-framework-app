package settings

import (
	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/viperEx"
)

const commandsSettingName = "commands"

// Fetch the setting
func (theSettings) Commands() []string {
	return viperEx.GetStringSliceOrDefault(commandsSettingName, []string{"example"})
}
