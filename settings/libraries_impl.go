// Generated 2021-02-24 16:31:35 by go-framework development-version
package settings

import (
	"github.com/AtricoSoftware/go-framework-app/requirements"
)

type Library struct {
	URL string
	Version string
}

func ParseLibrariesSetting(setting interface{}) map[string]string {
	results := make(map[string]string, 0)
	// Add standard requirements
	for _,lib := range requirements.Requirements {
		results[lib] = ""
	}
	// Read config requirements
	for _,lib := range setting.([]interface{}) {
		results[lib.(string)] = ""
	}
	return results
}
