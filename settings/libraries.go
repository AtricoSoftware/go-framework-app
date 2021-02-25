// Generated 2021-02-25 15:41:38 by go-framework development-version
package settings


import (
	"github.com/atrico-go/viperEx"
)

const librariesSettingName = "Libraries"

// Lazy value
var librariesSettingLazy = NewLazyStringSliceValue(func () []string { return viperEx.GetStringSlice(librariesSettingName) })

// Fetch the setting
func (theSettings) Libraries() []string {
	return librariesSettingLazy.GetValue()
}
