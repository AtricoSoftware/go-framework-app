// Generated 2021-02-25 11:57:44 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
)

const librariesSettingName = "Libraries"

// Lazy value
var librariesSettingLazy struct {
	theValue []string
	hasValue bool
}

// Fetch the setting
func (theSettings) Libraries() []string {
	if !librariesSettingLazy.hasValue {
		librariesSettingLazy.theValue = viperEx.GetStringSlice(librariesSettingName)
		librariesSettingLazy.hasValue = true
	}
	return librariesSettingLazy.theValue
}
