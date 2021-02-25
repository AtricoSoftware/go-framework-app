// Generated 2021-02-25 13:40:05 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
)

const librariesSettingName = "Libraries"


// Fetch the setting
func (theSettings) Libraries() []string {
	if !librariesSettingLazy.hasValue {
		librariesSettingLazy.theValue = viperEx.GetStringSlice(librariesSettingName)
		librariesSettingLazy.hasValue = true
	}
	return librariesSettingLazy.theValue
}
// Lazy value
var librariesSettingLazy struct {
	theValue []string
	hasValue bool
}
