package settings

import (
	"github.com/spf13/viper"

	"dev.azure.com/MAT-OCS/ConditionInsight/_git/ma.ci.go-framework-app/requirements"
)

const librariesSettingName = "libraries"

type Library struct {
	URL string
	Version string
}

// Fetch the setting
func (theSettings) Libraries() map[string]string {
	set := viper.Get(librariesSettingName)
	results := make(map[string]string, 0)
	// Add standard requirements
	for _,lib := range requirements.Requirements {
		results[lib] = ""
	}
	// Read config requirements
	for _,lib := range set.([]interface{}) {
		results[lib.(string)] = ""
	}
	return results
}
