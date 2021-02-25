// Generated 2021-02-25 10:44:36 by go-framework development-version
package settings

import (
	"github.com/atrico-go/viperEx"
)

const librariesSettingName = "Libraries"

// Fetch the setting
func (theSettings) Libraries() []string {
	return viperEx.GetStringSlice(librariesSettingName)
}
