// Generated 2021-03-04 17:50:38 by go-framework v1.6.0
// SECTION-START: Framework
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

// SECTION-END
