// Generated 2021-02-25 16:45:33 by go-framework v1.5.0
package settings

import (
	"github.com/atrico-go/viperEx"
)

const librariesSettingName = "Libraries"

// Lazy value
var librariesSettingLazy = NewLazyStringSliceValue(func() []string { return viperEx.GetStringSlice(librariesSettingName) })

// Fetch the setting
func (theSettings) Libraries() []string {
	return librariesSettingLazy.GetValue()
}
